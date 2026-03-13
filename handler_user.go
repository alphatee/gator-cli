package main 

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name) 
	}

	if err := s.cfg.SetUser(cmd.Args[0]); err != nil {
		return fmt.Errorf("Username does not exist: %s <name>", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}
