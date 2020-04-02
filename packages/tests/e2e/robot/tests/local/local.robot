*** Settings ***
Documentation       Check local training and packaging
Test Timeout        5 minutes
Resource            ../../resources/keywords.robot
Resource            ./resources/variables.robot
Library             Collections
Library             odahuflow.robot.libraries.utils.Utils
Library             odahuflow.robot.libraries.model.Model
#Suite Setup         Run Keywords
#...                 Set Environment Variable  ODAHUFLOW_CONFIG  ${LOCAL_CONFIG}  AND
#...                 Login to the api and edge  AND
#...                 Cleanup example resources  ${WINE_ID}
#Suite Teardown      Run Keywords
#...                 Cleanup example resources  ${WINE_ID}  AND
#...                 Remove file  ${LOCAL_CONFIG}
Force Tags          local e2e

*** Keywords ***
odahuflowctl local train run should pass
    [Documentation]  Verify variuos valid combination of model training ID, Path to a odahu-flow manifest file or
...                  Path to a dir with ODAHU-flow manifest files and dir where model artifacts will be saved
    [Arguments]      ${train id}
                     ${manifest file}
                     ${output dir}=
    ${resp}=         StrictShell  odahuflowctl local train run      ${train id}     ${manifest file}       ${output dir}

odahuflowctl local train run should fail
    [Documentation]  Verify variuos valid combination of model training ID, Path to a odahu-flow manifest file or
...                  Path to a dir with ODAHU-flow manifest files and dir where model artifacts will be saved
    [Arguments]      ${train id}
                     ${manifest file}
                     ${output dir}=
    ${resp}=         FailedShell  odahuflowctl local train run ${train id} ${manifest file} ${output dir}
                    #  Should contain   ${res.stderr}  ${train id}  ${file}  ${dir}  ${output dir}

*** Test Cases ***
    #odahuflowctl
    #    ${resp}=        StrictShell  odahuflowctl
    #                    Should contain  ${resp.stdout}  ${CTL_OUT}
    #
    #odahuflowctl --help
    #    ${resp}=        StrictShell  odahuflowctl --help
    #                    Should contain  ${resp.stdout}  ${CTL_OUT}
    #
    #odahuflowctl local
    #    ${resp}=        StrictShell  odahuflowctl local
    #                    Should be Equal  ${resp.stdout}  ${CTL_LOCAL_OUT}
    #
    #odahuflowctl local --help
    #    ${resp}=        StrictShell  odahuflowctl local --help
    #                    Should be Equal  ${resp.stdout}  ${CTL_LOCAL_OUT}
    #
    #odahuflowctl local training
    #    ${resp}=        StrictShell  odahuflowctl local training
    #                    Should be Equal  ${resp.stdout}  ${CTL_LOCAL_TRAINING_OUT}
    #odahuflowctl local training --help
    #    ${resp}=        StrictShell  odahuflowctl local training --help
    #                    Should be Equal  ${resp.stdout}  ${CTL_LOCAL_TRAINING_OUT}
    #
    #odahuflowctl local train
    #    ${resp}=        StrictShell  odahuflowctl local train
    #                    Should be Equal  ${resp.stdout}  ${CTL_LOCAL_TRAIN_OUT}
    #
    #odahuflowctl local train --help
    #    ${resp}=        StrictShell  odahuflowctl local train --help
    #                    Should be Equal  ${resp.stdout}  ${CTL_LOCAL_TRAIN_OUT}
    #
    #odahuflowctl local train run
    #    ${resp}=        FailedShell  odahuflowctl local train run
    #                    Should be Equal  ${resp.stderr}  ${CTL_LOCAL_TRAIN_RUN_ERR}
    #
    #odahuflowctl local train run --help
    #    ${resp}=        StrictShell  odahuflowctl local train run --help
    #                    Should be Equal  ${resp.stdout}  ${CTL_LOCAL_TRAIN_RUN_OUT}

odahuflowctl local train run valid
    [Template]      odahuflowctl local train run should pass
    --id wine-id-train          -d resources/local/

odahuflowctl local train run invalid
    [Template]      odahuflowctl local train run should fail

