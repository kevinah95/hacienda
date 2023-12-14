## hacienda completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(hacienda completion bash)

To load completions for every new session, execute once:

#### Linux:

	hacienda completion bash > /etc/bash_completion.d/hacienda

#### macOS:

	hacienda completion bash > $(brew --prefix)/etc/bash_completion.d/hacienda

You will need to start a new shell for this setup to take effect.


```
hacienda completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [hacienda completion](/cmd/hacienda_completion/)	 - Generate the autocompletion script for the specified shell

