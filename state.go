package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/saniak-hub/gator/internal/config"
	"github.com/saniak-hub/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("the login handler expects a single argument, the username")
	}

	userName := cmd.args[0]
	s.config.SetUser(userName)
	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		fmt.Println("Error login")
		os.Exit(1)
	}
	fmt.Printf("%s has been set\n", userName)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("the register handler require one argument, name")
	}
	new_user := database.CreateUserParams{
		ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.args[0],
	}

	user, err := s.db.CreateUser(context.Background(), new_user)
	if err != nil {
		return err
		os.Exit(1)
	}

	s.config.SetUser(user.Name)
	fmt.Printf("%s was created\n", user.Name)
	return nil
}

func handlerReset(s *state, cmd command) error {
	if err := s.db.TruncateAllData(context.Background()); err != nil {
		return err
	}
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	current := config.Read().CurrentUserName

	for _, user := range users {
		if user == current {
			fmt.Printf("%s (current)\n", user)
			continue
		}
		fmt.Println(user)
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("agg require one argument time between requests")
	}

	duration, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Invalide duration: %w", err)
	}

	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
	return nil
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return errors.New("Add feed require two arguments name and url")
	}

	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	}
	createdFeed, err := s.db.CreateFeed(context.Background(), feed)
	if err != nil {
		return err
	}

	fmt.Println(createdFeed)

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    createdFeed.ID,
	}
	_, err = s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Printf("Name: %s url: %s User: %s\n", feed.Name, feed.Url, feed.Name_2)
	}
	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("Requires one argument url")
	}

	feed, err := s.db.GetFeedWithUrl(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	ctx := context.Background()

	feedFollow, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return err
	}

	for _, feed := range feedFollow {
		fmt.Println(feed.FeedName)
	}
	return nil
}

func handlerUnFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("unfollow require one argument, url")
	}

	feed, err := s.db.GetFeedWithUrl(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	deleteFeedFollowParams := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	if err := s.db.DeleteFeedFollow(context.Background(), deleteFeedFollowParams); err != nil {
		return err
	}
	return nil
}

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32
	var ctx = context.Background()
	if len(cmd.args) == 1 {
		parselimit, err := strconv.ParseInt(cmd.args[0], 10, 32)
		if err != nil {
			return err
		}
		limit = int32(parselimit)
	} else {
		limit = 2
	}

	postParams := database.GetUserPostsParams{
		ID:    user.ID,
		Limit: limit,
	}
	posts, err := s.db.GetUserPosts(ctx, postParams)
	if err != nil {
		return err
	}
	for _, post := range posts {
		fmt.Printf("Title: %s  URL: %s\n", post.Title, post.Url)
	}
	return nil
}
