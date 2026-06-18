package main 

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
    		return fmt.Errorf("the login command expects a single argument: the username")
	}

	if err := s.cfg.SetUser(cmd.args[0]); err != nil {
		return fmt.Errorf("couldn't set user: %w", err)
	}

	fmt.Printf("User %s switched successfully!\n", cmd.args[0])
	return nil
}
