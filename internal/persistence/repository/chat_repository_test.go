package repository

import (
    "errors"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "symphony-api/internal/persistence/model"
)

func TestPutChat_Success(t *testing.T) {
    mockConn := new(MockPostgreConnection)
    repo := NewChatRepository(mockConn)

    input := &model.Chat{
        ChatId:    1,
        CreatedAt: time.Now(),
    }

    mockConn.On("Put", input.ToMap(), CHAT_TABLE_NAME).Return(nil)

    err := repo.Put(input)

    assert.NoError(t, err)
    mockConn.AssertExpectations(t)
}

func TestPutChat_Failure(t *testing.T) {
    mockConn := new(MockPostgreConnection)
    repo := NewChatRepository(mockConn)

    input := &model.Chat{
        ChatId:    1,
        CreatedAt: time.Now(),
    }

    mockConn.On("Put", input.ToMap(), CHAT_TABLE_NAME).Return(errors.New("db error"))

    err := repo.Put(input)

    assert.Error(t, err)
    mockConn.AssertExpectations(t)
}

func TestGetByChatId_Success(t *testing.T) {
    mockConn := new(MockPostgreConnection)
    repo := NewChatRepository(mockConn)

    constraint := map[string]any{
        "chat_id": int32(1),
    }

    dbResult := []map[string]any{
        {
            "chat_id":    int32(1),
            "created_at": time.Now(),
        },
    }

    mockConn.On("Get", constraint, CHAT_TABLE_NAME).Return(dbResult, nil)

    result, err := repo.GetByChatId(1)

    assert.NoError(t, err)
    assert.Equal(t, int32(1), result.ChatId)
    mockConn.AssertExpectations(t)
}

func TestGetByChatId_NotFound(t *testing.T) {
    mockConn := new(MockPostgreConnection)
    repo := NewChatRepository(mockConn)

    mockConn.On("Get", mock.Anything, CHAT_TABLE_NAME).Return([]map[string]any{}, nil)

    result, err := repo.GetByChatId(999)

    assert.Error(t, err)
    assert.Nil(t, result)
    mockConn.AssertExpectations(t)
}

func TestAddUserToChat_Success(t *testing.T) {
    mockConn := new(MockPostgreConnection)
    repo := NewChatRepository(mockConn)

    user := &model.User{UserId: 2}
    chat := &model.Chat{ChatId: 1}

    expectedMap := map[string]any{
        "chat_id": chat.ChatId,
        "user_id": user.UserId,
    }

    mockConn.On("Put", expectedMap, USER_TO_CHAT_TABLE).Return(nil)

    err := repo.AddUserToChat(user, chat)
    assert.NoError(t, err)
    mockConn.AssertExpectations(t)
}

func TestAddUserToChat_Failure(t *testing.T) {
    mockConn := new(MockPostgreConnection)
    repo := NewChatRepository(mockConn)

    user := &model.User{UserId: 2}
    chat := &model.Chat{ChatId: 1}

    expectedMap := map[string]any{
        "chat_id": chat.ChatId,
        "user_id": user.UserId,
    }

    mockConn.On("Put", expectedMap, USER_TO_CHAT_TABLE).Return(errors.New("db error"))

    err := repo.AddUserToChat(user, chat)
    assert.Error(t, err)
    mockConn.AssertExpectations(t)
}