package handlers

import (
	repository "database-example/repo"

	events "github.com/SOA-Tim-5/common/common/saga/complete_encounter"
	saga "github.com/SOA-Tim-5/common/common/saga/messaging"
)

type CompleteEncounterCommandHandler struct {
	followerRepo      *repository.FollowerRepository
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewCompleteEncounterCommandHandler(followerRepo *repository.FollowerRepository, publisher saga.Publisher, subscriber saga.Subscriber) (*CompleteEncounterCommandHandler, error) {
	o := &CompleteEncounterCommandHandler{
		followerRepo:      followerRepo,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	println("new complete encounter command handler")
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *CompleteEncounterCommandHandler) handle(command *events.UpdateLevelCommand) {
	id := (command.UpdateLevel.UserId)
	if id == "" {
		return
	}
	println("handle")
	reply := events.CompleteEncounterReply{UpdateLevel: command.UpdateLevel}

	switch command.Type {
	case events.UpdateFollower:
		oldUser, r1 := handler.followerRepo.GetUserById(command.UpdateLevel.UserId)
		if r1 == nil {
			println(oldUser.Username)
			println("Old level is ")
			print(oldUser.Level)
		}

		err := handler.followerRepo.UpdateUserLevelById(command.UpdateLevel.UserId, command.UpdateLevel.Level)
		if err != nil {
			reply.Type = events.FollowerUpdated
		} else {
			reply.Type = events.FollowerNotUpdated
		}
		newUser, r2 := handler.followerRepo.GetUserById(command.UpdateLevel.UserId)
		if r2 == nil {
			println(newUser.Username)
			println("New level is ")
			print(newUser.Level)
		}
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
