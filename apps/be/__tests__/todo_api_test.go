package __tests__

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/vizitiuRoman/hyf/__tests__/helper"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/config"
	"github.com/vizitiuRoman/hyf/internal/common/adapter/log"
	"github.com/vizitiuRoman/hyf/internal/domain/model"
	hyfv1 "github.com/vizitiuRoman/hyf/pkg/adapter/hyf/v1"
	pb "github.com/vizitiuRoman/hyf/pkg/gen/hyf/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestTodoAPISuite(t *testing.T) {
	suite.Run(
		t,
		helper.MustSuiteWithOptions(t, new(TodoAPISuite), helper.WithConfig(), helper.WithLogger()),
	)
}

type TodoAPISuite struct {
	suite.Suite
	ctx    context.Context
	cancel context.CancelFunc

	todoSVCClient pb.TodoSVCClient

	Cfg    *config.Config
	Logger log.Logger

	// suite
	mockTodo *model.Todo

	bearer string
}

func (s *TodoAPISuite) SetupSuite() {
	s.bearer = "omg"

	var err error
	s.ctx, s.cancel = context.WithCancel(context.Background())

	md := metadata.New(map[string]string{"authorization": "Bearer " + s.bearer})
	s.ctx = metadata.NewOutgoingContext(s.ctx, md)

	s.todoSVCClient, err = hyfv1.NewTodoSVCClient(
		s.ctx,
		s.Logger,
		&hyfv1.Config{
			Host:     "localhost",
			GrpcPort: 32772,
		},
		[]grpc.DialOption{}, nil,
	)
	s.Nil(err)

	s.mockTodo = &model.Todo{
		Title:       "name 2",
		Description: "desc 3",
	}
}

func (s *TodoAPISuite) TearDownSuite() {
	s.cancel()
}

func (s *TodoAPISuite) TestCreateTodo() {
	todo := pb.Todo{
		Title:       s.mockTodo.Title,
		Description: s.mockTodo.Description,
	}

	out, err := s.todoSVCClient.CreateTodo(s.ctx, &pb.CreateTodoIn{
		Todo: &todo,
	})
	s.Nil(err)

	s.NotNil(out.Todo)
	s.NotNil(out.Todo.Id)

	s.Equal(out.Todo.Title, todo.Title)
	s.Equal(out.Todo.Description, todo.Description)

	// set created id
	s.mockTodo.ID = out.Todo.Id
}

func (s *TodoAPISuite) TestUpdateTodo() {
	s.mockTodo.Title = "updated"

	s.Run("should successfully updated todo", func() {
		todo := pb.Todo{
			Id:          s.mockTodo.ID,
			Title:       s.mockTodo.Title,
			Description: s.mockTodo.Description,
		}

		out, err := s.todoSVCClient.UpdateTodo(s.ctx, &pb.UpdateTodoIn{
			Todo: &todo,
		})
		s.Nil(err)

		s.NotNil(out.Todo)
		s.NotNil(out.Todo.Id)

		s.Equal(out.Todo.Title, todo.Title)
		s.Equal(out.Todo.Description, todo.Description)
	})

	s.Run("should not found todo to update", func() {
		todo := pb.Todo{
			Id:          999999999,
			Title:       s.mockTodo.Title,
			Description: s.mockTodo.Description,
		}

		out, err := s.todoSVCClient.UpdateTodo(s.ctx, &pb.UpdateTodoIn{
			Todo: &todo,
		})
		s.NotNil(err)

		s.Nil(out)

		s.Equal(err.Error(), "rpc error: code = NotFound desc = cannot update todo: model was not found")
	})
}

func (s *TodoAPISuite) TestGetTodo() {
	s.Run("should found todo by id", func() {
		out, err := s.todoSVCClient.GetTodo(s.ctx, &pb.GetTodoIn{
			Id: s.mockTodo.ID,
		})

		s.Nil(err)

		s.NotNil(out.Todo)

		s.Equal(out.Todo.Title, s.mockTodo.Title)
		s.Equal(out.Todo.Description, s.mockTodo.Description)
		s.Equal(out.Todo.Id, s.mockTodo.ID)
	})

	s.Run("should not found todo", func() {
		out, err := s.todoSVCClient.GetTodo(s.ctx, &pb.GetTodoIn{
			Id: 999999999,
		})

		s.Nil(out)
		s.NotNil(err)
		s.Equal(err.Error(), "rpc error: code = NotFound desc = cannot find todo: model was not found")
	})
}

func (s *TodoAPISuite) TestGetTodos() {
	out, err := s.todoSVCClient.GetTodos(s.ctx, &pb.GetTodosIn{})
	s.Nil(err)
	s.NotZero(len(out.Todos))
}
