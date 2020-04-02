*** Variables ***
${CTL_OUT}                     local${SPACE * 18}Train and package locally
${CTL_LOCAL_OUT}               SEPARATOR=
...                            Usage: odahuflowctl local [OPTIONS] COMMAND [ARGS]...                     \n
...                                                                                                      \n
...                            \ \ Train and package locally                                             \n
...                                                                                                      \n
...                            Options:                                                                  \n
...                            \ \ --help \ Show this message and exit.                                  \n
...                                                                                                      \n
...                            Commands:                                                                 \n
...                            \ \ packaging \ Local packaging process.                                  \n
...                            \ \ training \ \ Local training process.
${CTL_LOCAL_TRAINING_OUT}      SEPARATOR=
...                            Usage: odahuflowctl local training [OPTIONS] COMMAND [ARGS]...            \n
...                                                                                                      \n
...                            \ \ Local training process.                                               \n
...                                                                                                      \n
...                            \ \ Alias for the command is train.                                       \n
...                                                                                                      \n
...                            Options:                                                                  \n
...                            \ \ --url TEXT${SPACE * 4}API server host                                 \n
...                            \ \ --token TEXT${SPACE * 2}API server jwt token                          \n
...                            \ \ --help${SPACE * 8}Show this message and exit.                         \n
...                                                                                                      \n
...                            Commands:                                                                 \n
...                            \ \ cleanup-artifacts${SPACE * 3}Delete all training local artifacts.     \n
...                            \ \ cleanup-containers${SPACE * 2}Delete all training docker containers.  \n
...                            \ \ list${SPACE * 16}Get list of local training artifacts.                \n
...                            \ \ run${SPACE * 17}Start a training process locally.
${CTL_LOCAL_TRAIN_OUT}         SEPARATOR=
...                            Usage: odahuflowctl local train [OPTIONS] COMMAND [ARGS]...               \n
...                                                                                                      \n
...                            \ \ Local training process.                                               \n
...                                                                                                      \n
...                            \ \ Alias for the command is train.                                       \n
...                                                                                                      \n
...                            Options:                                                                  \n
...                            \ \ --url TEXT${SPACE * 4}API server host                                 \n
...                            \ \ --token TEXT${SPACE * 2}API server jwt token                          \n
...                            \ \ --help${SPACE * 8}Show this message and exit.                         \n
...                                                                                                      \n
...                            Commands:                                                                 \n
...                            \ \ cleanup-artifacts${SPACE * 3}Delete all training local artifacts.     \n
...                            \ \ cleanup-containers${SPACE * 2}Delete all training docker containers.  \n
...                            \ \ list${SPACE * 16}Get list of local training artifacts.                \n
...                            \ \ run${SPACE * 17}Start a training process locally.
${CTL_LOCAL_TRAIN_RUN_ERR}     SEPARATOR=
...                            Usage: odahuflowctl local train run [OPTIONS]                             \n
...                            Try 'odahuflowctl local train run --help' for help.                       \n
...                                                                                                      \n
...                            Error: Missing option '--train-id' / '--id'.
${CTL_LOCAL_TRAIN_RUN_OUT}     SEPARATOR=
...                            Usage: odahuflowctl local train run [OPTIONS]                             \n
...                                                                                                      \n
...                            \ \ Start a training process locally.                                     \n
...                                                                                                      \n
...                            \ \ Usage example:                                                        \n
...                                                                                                      \n
...                            ${SPACE * 6}* odahuflowctl local train run --id examples-git              \n
...                                                                                                      \n

Usage example:

      * odahuflowctl local train run --id examples-git

Options:
  --train-id, --id TEXT        Model training ID  [required]
  -f, --manifest-file PATH     Path to a ODAHU-flow manifest file
  -d, --manifest-dir PATH      Path to a directory with ODAHU-flow manifest
                               files

  --output-dir, --output PATH  Directory where model artifact will be saved
  --help                       Show this message and exit.
