// This file is part of the JUSTtheTalkAPI distribution (https://github.com/jdudmesh/justthetalk-api).
// Copyright (c) 2021 John Dudmesh.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3.

// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package businesslogic

import (
	"justthetalk/connections"
	"justthetalk/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestFetchBlockedUsers(t *testing.T) {

	discussionId := uint(25876)

	folderCache := NewFolderCache()
	discussionCache := NewDiscussionCache(folderCache)

	discussion := discussionCache.UnsafeGet(discussionId)
	blockedUsers := discussionCache.BlockedUsers(discussion)
	if _, exists := blockedUsers[uint(2994)]; !exists {
		t.Fail()
	}

}

func TestBlockUser(t *testing.T) {

	discussionId := uint(25876)

	userCache := NewUserCache()
	folderCache := NewFolderCache()
	discussionCache := NewDiscussionCache(folderCache)

	targetUser := userCache.Get(5540)
	adminUser := userCache.Get(50)
	discussion := discussionCache.UnsafeGet(discussionId)
	blockedUsers := discussionCache.BlockOrUnblockUser(discussion, targetUser, true, adminUser)
	if _, exists := blockedUsers[uint(5540)]; !exists {
		t.Fail()
	}

}

func TestUnblockUser(t *testing.T) {

	discussionId := uint(25876)

	userCache := NewUserCache()
	folderCache := NewFolderCache()
	discussionCache := NewDiscussionCache(folderCache)

	targetUser := userCache.Get(5540)
	adminUser := userCache.Get(50)
	discussion := discussionCache.UnsafeGet(discussionId)
	blockedUsers := discussionCache.BlockOrUnblockUser(discussion, targetUser, false, adminUser)
	if _, exists := blockedUsers[uint(5540)]; exists {
		t.Fail()
	}

}

func TestAdminDeleteUndeletePost(t *testing.T) {
	t.Fail()
}

func TestAdminGetReports(t *testing.T) {
	t.Fail()
}

func TestAdminCreateAndGetComments(t *testing.T) {
	t.Fail()
}

func TestLockUnlockDiscussion(t *testing.T) {
	t.Fail()
}

func TestAdminPremodDiscussion(t *testing.T) {
	t.Fail()
}

func TestAdminDeleteDiscussion(t *testing.T) {
	t.Fail()
}

func TestAdminMoveDiscussion(t *testing.T) {
	t.Fail()
}

func TestAdminEraseDiscussion(t *testing.T) {
	t.Fail()
}

func TestModerationQueue(t *testing.T) {
	// TODO - clear queue, create reports
	connections.WithDatabase(60*time.Second, func(db *gorm.DB) {

		folderCache := NewFolderCache()
		discussionCache := NewDiscussionCache(folderCache)

		posts := GetModerationQueue(folderCache, discussionCache, db)
		if len(posts) == 0 {
			t.Error("No posts")
		}
	})
}

func TestSearchUsers(t *testing.T) {

	connections.WithDatabase(60*time.Second, func(db *gorm.DB) {

		results := SearchUsers("johnny", db)
		if len(results) == 0 {
			t.Error("No results")
		}

		results = SearchUsers("@@@", db)
		if len(results) > 0 {
			t.Error("Unexpected results")
		}

	})
}

func TestSetUserStatus(t *testing.T) {

	connections.WithDatabase(60*time.Second, func(db *gorm.DB) {

		userCache := NewUserCache()
		targetUser := userCache.Get(5540)
		adminUser := userCache.Get(50)

		fieldMap := make(map[string]interface{})

		var updated *model.User
		var err error

		fieldMap["isWatch"] = false
		updated, err = SetUserStatus(targetUser, fieldMap, adminUser, userCache, db)
		assert.NoError(t, err)
		if updated.IsWatch {
			t.Error("Unexpected: isWatch")
		}
		fieldMap["isWatch"] = true
		updated, err = SetUserStatus(targetUser, fieldMap, adminUser, userCache, db)
		assert.NoError(t, err)
		if !updated.IsWatch {
			t.Error("Unexpected: isWatch")
		}

	})
}
