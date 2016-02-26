package models

import (
        mgo "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
        "code.google.com/p/go.crypto/bcrypt"
        log "github.com/Sirupsen/logrus"
        "time"
)

// User is an information for the stored User object
type User struct {
        ID bson.ObjectId `bson:"_id,omitempty"`
        Email string `bson:"e"`
        Username string `bson:"u"`
        Password []byte  `bson:"p"`
        Timestamp time.Time `bson:"t"`
}

// HashPassword is used to bcrypt/hash the actual string
func (user *User) HashPassword(password string) {
        hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        if err != nil {
                log.Fatalf("Couldn't hash password: %v", err)
                panic(err)
        }
        user.Password = hash
}

// GetUserByEmail is what it is
func GetUserByEmail(database *mgo.Database, email string) (user *User) {
        err := database.C("users").Find(bson.M{"e": email}).One(&user)

        if err != nil {
                log.Warningf("Can't get user by email: %v", err)
        }
        return
}

// InsertUser is inserting a new user to the db
func InsertUser(database *mgo.Database, user *User) error {
        user.ID = bson.NewObjectId()
        return database.C("users").Insert(user)
}
