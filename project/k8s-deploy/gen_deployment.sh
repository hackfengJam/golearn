#!/bin/bash

#set -ex

source ../_apps.sh
HUB_PREFIX=ptapp.cn:8443/pt-cluster
GIN_MODE_ENV=${GIN_MODE:=test}

PT_ENIGMA_URL=${ENIGMA_URL:=http://192.168.18.59:30090/keys}

for app in ${apps[@]}
do
    deployment_yml=_gen_"$app"_deploy.yml
    echo "Generating deployment" $app  "->" $deployment_yml
    HEALTHZ_PORT=`python find_service_healthz_port.py $app`
    sed  "s/{{service}}/$app/g" deployment.tmpl >  $deployment_yml
    sed -i "" 's|{{port}}|'"$HEALTHZ_PORT"'|g'  $deployment_yml   2> /dev/null  # 在 Mac 上 sed -i 需要加空字符串
    sed -i "" 's|{{hub_prefix}}|'"$HUB_PREFIX"'|g'  $deployment_yml  2> /dev/null
    sed -i "" 's|{{gin_mode}}|'"$GIN_MODE_ENV"'|g'  $deployment_yml  2> /dev/null
    sed -i "" 's|{{ENIGMA_URL}}|'"$PT_ENIGMA_URL"'|g'  $deployment_yml  2> /dev/null

    sed -i 's|{{port}}|'"$HEALTHZ_PORT"'|g'  $deployment_yml  2> /dev/null
    sed -i 's|{{hub_prefix}}|'"$HUB_PREFIX"'|g'  $deployment_yml 2> /dev/null
    sed -i 's|{{gin_mode}}|'"$GIN_MODE_ENV"'|g'  $deployment_yml 2> /dev/null
    sed -i 's|{{ENIGMA_URL}}|'"$PT_ENIGMA_URL"'|g'  $deployment_yml 2> /dev/null

done