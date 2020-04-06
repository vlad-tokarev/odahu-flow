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

