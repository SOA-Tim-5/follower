package handlers

import (
	"context"
	"database-example/model"
	repository "database-example/repo"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type KeyProduct struct{}

type FollowerHandler struct {
	logger *log.Logger
	repo   *repository.FollowerRepository
}

func NewFollowersHandler(l *log.Logger, r *repository.FollowerRepository) *FollowerHandler {
	return &FollowerHandler{l, r}
}

func (f *FollowerHandler) CreateUser(rw http.ResponseWriter, h *http.Request) {
	user := h.Context().Value(KeyProduct{}).(*model.User)
	userSaved, err := f.repo.SaveUser(user)
	if err != nil {
		f.logger.Print("Database exception: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if userSaved {
		f.logger.Print("New user saved to database")
		rw.WriteHeader(http.StatusCreated)
	} else {
		rw.WriteHeader(http.StatusConflict)
	}
}

func (f *FollowerHandler) GetUser(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	user, err := f.repo.ReadUser(id)
	if err != nil {
		f.logger.Print("Database exception: ", err)
	}

	err = user.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		f.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (f *FollowerHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		f.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}

func (f *FollowerHandler) MiddlewarePersonDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		user := &model.User{}
		err := user.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			f.logger.Fatal(err)
			return
		}
		ctx := context.WithValue(h.Context(), KeyProduct{}, user)
		h = h.WithContext(ctx)
		next.ServeHTTP(rw, h)
	})
}

func (f *FollowerHandler) CreateFollowing(rw http.ResponseWriter, h *http.Request) {
	newFollowing := h.Context().Value(KeyProduct{}).(*model.UserFollowing)
	user := model.User{}
	userToFollow := model.User{}
	user.Id = newFollowing.UserId
	user.Username = newFollowing.Username
	user.Image = newFollowing.Image
	userToFollow.Id = newFollowing.FollowingUserId
	userToFollow.Username = newFollowing.FollowingUsername
	userToFollow.Image = newFollowing.FollowingImage
	err := f.repo.SaveFollowing(&user, &userToFollow)
	if err != nil {
		f.logger.Print("Database exception: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	user = model.User{}
	jsonData, _ := json.Marshal(user)
	rw.Write(jsonData)
}
func (f *FollowerHandler) MiddlewareFollowingDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		newFollowing := &model.UserFollowing{}
		err := newFollowing.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			f.logger.Fatal(err)
			return
		}
		ctx := context.WithValue(h.Context(), KeyProduct{}, newFollowing)
		h = h.WithContext(ctx)
		next.ServeHTTP(rw, h)
	})
}
func (f *FollowerHandler) GetFollowings(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["userId"]
	users, err := f.repo.GetFollowings(id)
	if err != nil {
		f.logger.Print("Database exception: ", err)
	}
	if users == nil {
		users = model.Users{}
		jsonData, _ := json.Marshal(users)
		rw.Write(jsonData)
		return
	}
	err = users.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		f.logger.Fatal("Unable to convert to json :", err)
		return
	}
}
func (f *FollowerHandler) GetFollowers(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["userId"]
	users, err := f.repo.GetFollowers(id)
	if err != nil {
		f.logger.Print("Database exception: ", err)
	}
	if users == nil {
		users = model.Users{}
		jsonData, _ := json.Marshal(users)
		rw.Write(jsonData)
		return
	}
	err = users.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		f.logger.Fatal("Unable to convert to json :", err)
		return
	}
}
func (f *FollowerHandler) GetRecommendations(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["userId"]
	users, err := f.repo.GetRecommendations(id)
	if err != nil {
		f.logger.Print("Database exception: ", err)
	}
	if users == nil {
		users = model.Users{}
		jsonData, _ := json.Marshal(users)
		rw.Write(jsonData)
		return
	}
	err = users.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		f.logger.Fatal("Unable to convert to json :", err)
		return
	}
}
