package services

// func TestCreateBattle(t *testing.T) {
// 	repoMock := new(mocks.MockBattleRepository)

// 	arg := params.CreateAutoBattle{
// 		BattleName: "Battle 1",
// 		Pokemons:   []int{1, 2, 3, 4, 5},
// 	}

// 	listTests := []struct {
// 		name    string
// 		mock    func()
// 		args    params.CreateAutoBattle
// 		want    *models.Battle
// 		wantErr bool
// 	}{
// 		{
// 			name: "Create Battle",
// 			mock: func() {
// 				data := models.Battle{
// 					Name: arg.BattleName,
// 				}
// 				repoMock.On("CreateBattle", arg).Return(data).Once()
// 			},
// 			args: arg,
// 			want: &models.Battle{
// 				ID:   1,
// 				Name: "Battle 1",
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Create Battle",
// 			mock: func() {
// 				data := models.Battle{
// 					Name: arg.BattleName,
// 				}
// 				err := errors.New("error")
// 				repoMock.On("CreateBattle", data).Return(err).Once()
// 			},
// 			args:    arg,
// 			want:    &models.Battle{},
// 			wantErr: false,
// 		},
// 	}

// 	for _, test := range listTests {
// 		test.mock()

// 		battleService := BattleService{battleRepo: repoMock}

// 		res := battleService.CreateAutoBattle(test.args)

// 		if test.wantErr {
// 			assert.Equal(t, models.Battle{}, res)
// 		}

// 		if !test.wantErr {
// 			assert.Equal(t, res, test.want)
// 		}
// 	}
// }
