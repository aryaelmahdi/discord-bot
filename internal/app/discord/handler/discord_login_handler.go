package handler

// func (handler *DiscordHandlerImpl) Login(s *discordgo.Session, i *discordgo.InteractionCreate) error {
// 	username := i.ApplicationCommandData().Options[0].StringValue()
// 	password := i.ApplicationCommandData().Options[1].StringValue()

// 	if username == "" || password == "" {
// 		// Respond with an error message if parameters are missing
// 		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 			Type: discordgo.InteractionResponseChannelMessageWithSource,
// 			Data: &discordgo.InteractionResponseData{
// 				Content: "Please provide both username and password.",
// 			},
// 		})
// 		return errors.New("no password or username provided")
// 	}

// 	// Perform login logic (replace with your actual login logic)
// 	request := web.DiscordLogin{
// 		Username: username,
// 		Password: password,
// 	}
// 	err := handler.Service.Login(&request)
// 	if err != nil {
// 		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 			Type: discordgo.InteractionResponseChannelMessageWithSource,
// 			Data: &discordgo.InteractionResponseData{
// 				Content: fmt.Sprintf("Error logging in: %v", err),
// 			},
// 		})
// 		return err
// 	}

// 	// Respond with success message
// 	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 		Type: discordgo.InteractionResponseChannelMessageWithSource,
// 		Data: &discordgo.InteractionResponseData{
// 			Content: "Login successful!",
// 		},
// 	})
// 	return nil
// }
