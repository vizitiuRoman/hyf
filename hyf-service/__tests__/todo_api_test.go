package __tests__

//
//func TestTodoAPISuite(t *testing.T) {
//	suite.Run(
//		t,
//		helper.MustSuiteWithOptions(t, new(TodoAPISuite), helper.WithConfig(), helper.WithLogger()),
//	)
//}
//
//type TodoAPISuite struct {
//	suite.Suite
//	ctx    context.Context
//	cancel context.CancelFunc
//
//	todoSVCClient pb.TodoSVCClient
//
//	Cfg    *helper.Config
//	Logger logger.Logger
//
//	// suite
//	mockTodo *model.Todo
//
//	bearer string
//}
//
//func (s *TodoAPISuite) SetupSuite() {
//	var err error
//	s.ctx, s.cancel = context.WithCancel(context.Background())
//
//	authSVCClient, err := hyfv1.NewAuthSVCClient(
//		s.ctx,
//		s.Logger,
//		&hyfv1.Config{
//			Host:     s.Cfg.Grpc.Host,
//			GrpcPort: s.Cfg.Grpc.Port,
//		},
//		[]grpc.DialOption{}, nil,
//	)
//
//	out, err := authSVCClient.Login(s.ctx, &pb.LoginIn{
//		Email:    s.Cfg.User.Email,
//		Password: s.Cfg.User.Password,
//	})
//	s.Nil(err)
//
//	s.bearer = out.Token
//
//	md := metadata.New(map[string]string{"authorization": "Bearer " + s.bearer})
//	s.ctx = metadata.NewOutgoingContext(s.ctx, md)
//
//	s.todoSVCClient, err = hyfv1.NewTodoSVCClient(
//		s.ctx,
//		s.Logger,
//		&hyfv1.Config{
//			Host:     s.Cfg.Grpc.Host,
//			GrpcPort: s.Cfg.Grpc.Port,
//		},
//		[]grpc.DialOption{}, nil,
//	)
//	s.Nil(err)
//
//	s.mockTodo = &model.Todo{
//		Title:       "name 2",
//		Description: "desc 3",
//	}
//}
//
//func (s *TodoAPISuite) TearDownSuite() {
//	s.cancel()
//}
//
//func (s *TodoAPISuite) TestCreateTodo() {
//	todo := pb.Todo{
//		Title:       s.mockTodo.Title,
//		Description: s.mockTodo.Description,
//	}
//
//	out, err := s.todoSVCClient.CreateTodo(s.ctx, &pb.CreateTodoIn{
//		Todo: &todo,
//	})
//	s.Nil(err)
//
//	s.NotNil(out.Todo)
//	s.NotNil(out.Todo.Id)
//
//	s.Equal(out.Todo.Title, todo.Title)
//	s.Equal(out.Todo.Description, todo.Description)
//
//	// set created id
//	s.mockTodo.ID = out.Todo.Id
//}
//
//func (s *TodoAPISuite) TestUpdateTodo() {
//	s.mockTodo.Title = "updated"
//
//	s.Run("should successfully updated todo", func() {
//		todo := pb.Todo{
//			Id:          s.mockTodo.ID,
//			Title:       s.mockTodo.Title,
//			Description: s.mockTodo.Description,
//		}
//
//		out, err := s.todoSVCClient.UpdateTodo(s.ctx, &pb.UpdateTodoIn{
//			Todo: &todo,
//		})
//		s.Nil(err)
//
//		s.NotNil(out.Todo)
//		s.NotNil(out.Todo.Id)
//
//		s.Equal(out.Todo.Title, todo.Title)
//		s.Equal(out.Todo.Description, todo.Description)
//	})
//
//	s.Run("should not found todo to update", func() {
//		todo := pb.Todo{
//			Id:          999999999,
//			Title:       s.mockTodo.Title,
//			Description: s.mockTodo.Description,
//		}
//
//		out, err := s.todoSVCClient.UpdateTodo(s.ctx, &pb.UpdateTodoIn{
//			Todo: &todo,
//		})
//		s.NotNil(err)
//
//		s.Nil(out)
//
//		s.Equal(err.Error(), "rpc error: code = NotFound desc = cannot update todo: model was not found")
//	})
//}
//
//func (s *TodoAPISuite) TestGetTodo() {
//	s.Run("should found todo by id", func() {
//		out, err := s.todoSVCClient.GetTodo(s.ctx, &pb.GetTodoIn{
//			Id: s.mockTodo.ID,
//		})
//
//		s.Nil(err)
//
//		s.NotNil(out.Todo)
//
//		s.Equal(out.Todo.Title, s.mockTodo.Title)
//		s.Equal(out.Todo.Description, s.mockTodo.Description)
//		s.Equal(out.Todo.Id, s.mockTodo.ID)
//	})
//
//	s.Run("should not found todo", func() {
//		out, err := s.todoSVCClient.GetTodo(s.ctx, &pb.GetTodoIn{
//			Id: 999999999,
//		})
//
//		s.Nil(out)
//		s.NotNil(err)
//		s.Equal(err.Error(), "rpc error: code = NotFound desc = cannot find todo: model was not found")
//	})
//}
//
//func (s *TodoAPISuite) TestGetTodos() {
//	out, err := s.todoSVCClient.GetTodos(s.ctx, &pb.GetTodosIn{})
//	s.Nil(err)
//	s.NotZero(len(out.Todos))
//}
