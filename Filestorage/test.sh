#!/bin/bash

#defaults
argc=$# 

api=bridgeserver
latency=0
loop=0
host=$1

declare -a array=("KOR" "IND" "USA" "NLD" "CHN" "DEU" "GBR" "FRA")
now=$(date +"%T")
current_date_time="`date +%d-%m-%Y`";
filename="output_"$current_date_time"_"$now.txt

#Bridgeserver Api's
setupvpnURL=$1"/api/v1/partners/samsung/users/sessions"
changeplanURL=$1"/api/v1/partners/samsung/users/plans?action=changeplan"
checkuserquotaURL=$1"/api/v1/partners/samsung/users/usage"
planusedURL=$1"/api/v1/partners/samsung/users/usage?fields=isused"
setURL=$1"/api/v1/partners/samsung/users/filter"
getURL=$1"/api/v1/partners/samsung/users/filter?fields=status"



if [ $argc -lt 1 ]; then
        echo "[ERROR] not enough arguments"
                echo "pass the country Code as paramater \n Example \n "
				echo "For Example - ./latency.sh  https://vpnbridgeirlstg.mcafee.com KOR"
				
				
               
        exit
fi




for i in {0..7}
do

     
       
	   currentregion=${array[$i]}
	   serviceid="SID18-August-"
       serviceid+="$currentregion"
       pskid=$currentregion"1@gmail.com"
       

	   
	    
        
        echo "Setupvpn For " $currentregion
		
		 
        echo "================"
		echo "Date&Time :"$current_date_time"_"$now
		echo "Region" :$currentregion
        echo "Api    :" $setupvpnURL
		echo "Param  :" '{"service_id": "'$serviceid'","profile_version": "1.1","vpn_id": "'$pskid'","psk": "111000","ike_port": 35620,"natt_port": 35621,"client_ip": "211.220.194.0"}'
		setupvpnout=$(curl -X POST  $setupvpnURL  -H 'content-type: application/json' -d '{"service_id": "'$serviceid'","profile_version": "1.1","vpn_id": "sandy@gmail.com","psk": "100100","ike_port": 35620,"natt_port": 35621,"client_ip": "10.23.50.23"}')
        echo "Output: " $setupvpnout
        echo " "
        
       

        echo "Checkuserquota For "$currentregion
        echo "================"
		echo "Date&Time :" $current_date_time"_"$now
		echo "Region" :$currentregion
		echo "Api   :" $checkuserquotaURL
		echo "Param :" '{"service_id": "'$serviceid'","future_recurrence": "1"}'
		checkuserout=$(curl -X POST  $checkuserquotaURL -H 'content-type: application/json' -d '{"service_id": "'$serviceid'","future_recurrence": "1"}')
        echo "Output: " $checkuserout
         echo " "

    
done >> $filename    


