*** Settings ***
Documentation       Check local training and packaging
#  Test Timeout        5 minutes
Resource            ../../resources/keywords.robot
Resource            ./resources/variables.robot
Library             Collections
Library             odahuflow.robot.libraries.utils.Utils
Library             odahuflow.robot.libraries.model.Model
#  Suite Setup         Run Keywords
#  ...                 Set Environment Variable  ODAHUFLOW_CONFIG  ${LOCAL_CONFIG}  AND
#  ...                 Login to the api and edge  AND
#  ...                 Cleanup example resources  ${WINE_ID}
#  Suite Teardown      Run Keywords
#  ...                 Cleanup example resources  ${WINE_ID}  AND
#  ...                 Remove file  ${LOCAL_CONFIG}
#   Test Setup
#   Test Teardown
Force Tags          local e2e

*** Variables ***
@{TRAINING LIST}
${manigest dir}    resources/local

*** Keywords ***
odahuflowctl local train run should pass
    [Documentation]  Verify variuos valid combination of model training ID, Path to a odahu-flow manifest file or
    ...              Path to a dir with ODAHU-flow manifest files and dir where model artifacts will be saved
    [Arguments]      ${train id}  ${manifest file}  ${output dir}=
    ${resp}=         StrictShell  odahuflowctl --verbose local train run ${train id} ${manifest file} ${output dir}
                     Append To List  ${TRAINING LIST}  ${train id}
                     Log  ${resp.stdout}

odahuflowctl local train run should fail
    [Documentation]  Verify variuos invalid cases of training options
    [Arguments]      ${train id}  ${manifest file}  ${output dir}=
    ${resp}=         FailedShell  odahuflowctl local train run ${train id} ${manifest file} ${output dir}
                    #  Should contain   ${resp.stderr}  ${train id}  ${file}  ${dir}  ${output dir}

odahuflowctl local pack run should pass
    [Documentation]  Verify variuos valid combination of local model packaging options
    [Arguments]      ${pack id}  ${manifest file}  ${artifact name}  ${artifact path}=  ${disable-package-targets}=
    ${resp}=         StrictShell  odahuflowctl --verbose local pack run ${pack id} ${manifest file} ${artifact name} ${artifact path} ${disable-package-targets}

odahuflowctl local pack run should fail
    [Documentation]  Verify variuos invalid cases of training options
    [Arguments]      ${train id}  ${manifest file}  ${output dir}=
    ${resp}=         FailedShell  odahuflowctl local pack run ${train id} ${manifest file} ${output dir}

odahuflowctl test packaged model
    StrictShell  odahuflowctl --verbose dep create -f ${manifest dir}/deployment.odahuflow.yaml --image ${res.stdout} --id ${example_id}

    Wait Until Keyword Succeeds  1m  0 sec  StrictShell  odahuflowctl model info --mr ${example_id}
    Wait Until Keyword Succeeds  1m  0 sec  StrictShell  odahuflowctl model invoke --mr ${example_id} --json-file $${manifest dir}/request.json

    ${res}=  Shell  odahuflowctl model invoke --mr ${example_id} --json-file ${manifest dir}/request.json --jwt wrong-token
    should not be equal  ${res.rc}  0

*** Test Cases ***
#odahuflowctl
#    ${resp}=        StrictShell  odahuflowctl
#                    Should Contain  ${resp.stdout}  ${CTL_OUT}
#
#odahuflowctl --help
#    ${resp}=        StrictShell  odahuflowctl --help
#                    Should Contain  ${resp.stdout}  ${CTL_OUT}
#
#odahuflowctl local
#    ${resp}=        StrictShell  odahuflowctl local
#                    Should Be Equal  ${resp.stdout}  ${CTL_LOCAL_OUT}
#
#odahuflowctl local --help
#    ${resp}=        StrictShell  odahuflowctl local --help
#                    Should Be Equal  ${resp.stdout}  ${CTL_LOCAL_OUT}
#
#odahuflowctl local training
#    ${resp}=        StrictShell  odahuflowctl local training
#                    Should Be Equal  ${resp.stdout}  ${CTL_LOCAL_TRAINING_OUT}
#odahuflowctl local training --help
#    ${resp}=        StrictShell  odahuflowctl local training --help
#                    Should Be Equal  ${resp.stdout}  ${CTL_LOCAL_TRAINING_OUT}
#
#odahuflowctl local train
#    ${resp}=        StrictShell  odahuflowctl local train
#                    Should Be Equal  ${resp.stdout}  ${CTL_LOCAL_TRAIN_OUT}
#
#odahuflowctl local train --help
#    ${resp}=        StrictShell  odahuflowctl local train --help
#                    Should Be Equal  ${resp.stdout}  ${CTL_LOCAL_TRAIN_OUT}
#
#odahuflowctl local train run
#    ${resp}=        FailedShell  odahuflowctl local train run
#                    Should Be Equal  ${resp.stderr}  ${CTL_LOCAL_TRAIN_RUN_ERR}
#

#--------- TRAINING RUN -----------
odahuflowctl local train run valid
    [Template]      odahuflowctl local train run should pass
    #  ${train id}                                 ${manifest file}                                         ${output dir}
    --id wine-id-d                                -d ${manifest dir}/
    --id wine-id-d-output                         -d ${manifest dir}/                                      --output ${CURDIR}/outputs/
#    --id wine-id-d-ouput-dir                      -d ${manifest dir}/                                      --output-dir ${CURDIR}/outputs/
#    --id wine-id-f                                -f ${manifest dir}/training.odahu.yaml
#    --id wine-id-f-output                         -f ${manifest dir}/training.odahu.yaml                   --output ${CURDIR}/outputs/
#    --id wine-id-f-output-dir                     -f ${manifest dir}/training.odahu.yaml                   --output-dir ${CURDIR}/outputs/
#    --id wine-id-dir                              --manifest-dir ${manifest dir}/
#    --id wine-id-dir-output                       --manifest-dir ${manifest dir}/                          --output ${CURDIR}/outputs/
#    --id wine-id-dir-output-dir                   --manifest-dir ${manifest dir}/                          --output-dir ${CURDIR}/outputs/
#    --id wine-id-file                             --manifest-file ${manifest dir}/training.odahu.yaml
#    --id wine-id-file-output                      --manifest-file ${manifest dir}/training.odahu.yaml      --output ${CURDIR}/outputs/
#    --id wine-id-file-output-dir                  --manifest-file ${manifest dir}/training.odahu.yaml      --output-dir ${CURDIR}/outputs/
#    --train-id wine-train-id-d                    -d ${manifest dir}/
#    --train-id wine-train-id-d-output             -d ${manifest dir}/                                      --output ${CURDIR}/outputs/
#    --train-id wine-train-id-d-ouput-dir          -d ${manifest dir}/                                      --output-dir ${CURDIR}/outputs/
#    --train-id wine-train-id-f                    -f ${manifest dir}/training.odahu.yaml
#    --train-id wine-train-id-f-output             -f ${manifest dir}/training.odahu.yaml                   --output ${CURDIR}/outputs/
#    --train-id wine-train-id-f-output-dir         -f ${manifest dir}/training.odahu.yaml                   --output-dir ${CURDIR}/outputs/
#    --train-id wine-train-id-dir                  --manifest-dir ${manifest dir}/
#    --train-id wine-train-id-dir-output           --manifest-dir ${manifest dir}/                          --output ${CURDIR}/outputs/
#    --train-id wine-train-id-dir-output-dir       --manifest-dir ${manifest dir}/                          --output-dir ${CURDIR}/outputs/
#    --train-id wine-train-id-file                 --manifest-file ${manifest dir}/training.odahu.yaml
#    --train-id wine-train-id-file-output          --manifest-file ${manifest dir}/training.odahu.yaml      --output ${CURDIR}/outputs/
#    --train-id wine-train-id-file-output-dir      --manifest-file ${manifest dir}/training.odahu.yaml      --output-dir ${CURDIR}/outputs/

odahuflowctl local train run invalid
    [Template]      odahuflowctl local train run should fail
    #  ${train id}                ${manifest file}                                         ${output dir}
    #  wrong train id
    --id wine-id-trai             --manifest-file ${manifest dir}/training.odahu.yaml      -output ${CURDIR}/outputs/
    #  manifest dir without training-id
    --train-id wine-id-d          --manifest-dir resources/
    #  manifest file without training-id
    --id wine-id-f-output         --manifest-file ${manifest dir}/packaging.odahu.yaml
    #  path to manifest file instead of manifest dir
    --id wine-id-trian            --manifest-dir ${manifest dir}/training.odahu.yaml
    #  path to manifest dir instead of manifest file
    --id wine-id-trian            --manifest-file ${manifest dir}/                         -output ${CURDIR}/outputs/

odahuflowctl local train list
    ${resp}=         StrictShell  odahuflowctl --verbose local train list
                     Should Contain  ${resp.stdout}  @{TRAINING LIST}


#odahuflowctl local pack run valid
#    [Template]     odahuflowctl local pack run should pass
#    #  ${pack id}                    ${manifest file}                                       ${artifact name}                   ${artifact path}                         ${disable-package-targets}
#    --id pack-id-d                   -d ${manifest dir}/                                    -a @{TRAINING LIST}
#    --id pack-id-d-output            -d ${manifest dir}/                                                                       --artifact-path ${CURDIR}/outputs/
#    --id pack-id-d-ouput-dir         -d ${manifest dir}/                                                                       --artifact-path ${CURDIR}/outputs/
#    --id pack-id-f                   -f ${manifest dir}/packaging.odahu.yaml
#    --id pack-id-f-output            -f ${manifest dir}/packaging.odahu.yaml                                                   --artifact-path ${CURDIR}/outputs/
#    --id pack-id-f-output-dir        -f ${manifest dir}/packaging.odahu.yaml                                                   --artifact-path ${CURDIR}/outputs/
#    --id pack-id-dir                 --manifest-dir ${manifest dir}/
#    --id pack-id-dir-output          --manifest-dir ${manifest dir}/                                                           --artifact-path ${CURDIR}/outputs/
#    --id pack-id-dir-output-dir      --manifest-dir ${manifest dir}/                                                           --artifact-path ${CURDIR}/outputs/
#    --id pack-id-file                --manifest-file ${manifest dir}/packaging.odahu.yaml
#    --id pack-id-file-output         --manifest-file ${manifest dir}/packaging.odahu.yaml                                      --artifact-path ${CURDIR}/outputs/
#    --id pack-id-file-output-dir     --manifest-file ${manifest dir}/packaging.odahu.yaml                                      --artifact-path ${CURDIR}/outputs/
#    --pack-id pack-id-d              -d ${manifest dir}/

odahuflowctl local pack run invalid
    #  ${pack id}                    ${manifest file}                                         ${artifact name}                   ${artifact path}         ${disable-package-targets}
    #  wrong pack id
    --id pack-id-invalid             --manifest-file ${manifest dir}/packaging.odahu.yaml     -output ${CURDIR}/outputs/
    #  manifest dir without pack-id
    --id pack-id-d                   --manifest-dir resources/
    #  manifest file without pack-id
    --id pack-id-f-output            --manifest-file ${manifest dir}/training.odahu.yaml
    #  path to manifest file instead of manifest dir
    --id pack-id-f                   --manifest-dir ${manifest dir}/packaging.odahu.yaml
    #  path to manifest dir instead of manifest file
    --id pack-id-dir                 --manifest-file ${manifest dir}/                         -output ${CURDIR}/outputs/
    # artifact path to the folder where artifact training doesn't exist
    --id pack-id-d-output            -d ${manifest dir}/               ${artifact path}=--artifact-path ${CURDIR}/



