mega-gochecker
====

### what is that

GoLang  MegaCLI parser 

### What you have

There are two programm with which you can grep info from megacli util, one you can use with zabbix dicovery rules

*megars* - megacli raid status, return status of raid: Optimal or Degraded  - this state from output of this command:

`megacli -LDInfo -Lall -aALL | grep State`

 *megads* - megacli disk status, return count of Media or Other error, command
 
` megacli -PDList -aALL | Enclosure Device ID:|Slot Number:|Inquiry Data:|Error Count:|state`

*megadcvr* - megacli discovery, return count of slots involved in raid

Example:

      #./megadcvr

     { 
	     "data":[
		    {
		    	"{#SLOTNUMBER}":"0"},
		    {
		    	"{#SLOTNUMBER}":"1"},
	    	{
		    	"{#SLOTNUMBER}":"2"},
	    	{
		    	"{#SLOTNUMBER}":"3"}]}

megadcvr - giving output which you can use with zabbix discovery rules

