package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

//go:generate go run "github.com/Jumpaku/cyamli/cmd/cyamli@latest" golang -schema-path=cli.yaml -out-path=cli.gen.go

func main() {
	cli := NewCLI()
	cli.FUNC = func(subcommand []string, input CLI_Input, inputErr error) (err error) {
		fmt.Println(cli.DESC_Detail())
		return nil
	}

	cli.Tags.FUNC = func(subcommand []string, input CLI_Tags_Input, inputErr error) (err error) {
		if inputErr != nil {
			fmt.Println(cli.Tags.DESC_Detail())
			return fmt.Errorf("invalid input: %w", inputErr)
		}

		url := `https://api.github.com/repos/Jumpaku/tz-offset-transitions/tags`
		resp, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("failed to get tags %q: %w", url, err)
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to get tags %q with status %q: %w", url, resp.Status, err)
		}

		var body []struct {
			Name string `json:"name"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			return fmt.Errorf("failed to extract tag names from response body: %w", err)
		}

		for _, t := range body {
			fmt.Println(t.Name)
		}

		return nil
	}

	cli.Download.FUNC = func(subcommand []string, input CLI_Download_Input, inputErr error) (err error) {
		if inputErr != nil {
			fmt.Println(cli.Download.DESC_Detail())
			return fmt.Errorf("invalid input: %w", inputErr)
		}

		out := os.Stdout
		if input.Opt_OutPath != "" {
			f, err := os.Create(input.Opt_OutPath)
			if err != nil {
				return fmt.Errorf("failed to create file %q: %w", input.Opt_OutPath, err)
			}

			defer f.Close()
			out = f
		}

		tzotVersion := input.Opt_Tag
		if tzotVersion == "" {
			url := `https://api.github.com/repos/Jumpaku/tz-offset-transitions/releases/latest`
			resp, err := http.Get(url)
			if err != nil {
				return fmt.Errorf("failed to get latest release %q: %w", url, err)
			}
			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("failed to get latest release %q with status %q: %w", url, resp.Status, err)
			}

			var body struct {
				TagName string `json:"tag_name"`
			}
			if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
				return fmt.Errorf("failed to extract tag_name from response body: %w", err)
			}

			tzotVersion = body.TagName
		}

		url := fmt.Sprintf(`https://github.com/Jumpaku/tz-offset-transitions/raw/%s/gen/tzot.json`, tzotVersion)
		resp, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("failed to get data %q: %w", url, err)
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to get data %q with status %q: %w", url, resp.Status, err)
		}

		if _, err := io.Copy(out, resp.Body); err != nil {
			return fmt.Errorf("failed to save data: %w", err)
		}

		return nil
	}

	if err := Run(cli, os.Args); err != nil {
		log.Panic(err)
	}
}
