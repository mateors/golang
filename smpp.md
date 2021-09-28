# SMPP Protocol

nc -z -v -u 10.10.8.8 20-80 2>&1 | grep succeeded

sudo nmap -sU -p- 10.10.8.8

https://linuxize.com/post/check-open-ports-linux/

SMPP is an asynchronous protocol. This means that an ESME or MC can send several
requests at a time to the other party.

SMPP operations are based on the exchange of operation PDUs between ESME and MC

Most SMPP sessions will be ESME-initiated, but as we have seen from Outbind, the MC can
itself be in a position to connect to an ESME that is configured for Outbind.


first establishing a secure and authenticated connection before commencing the
SMPP session. 


## interface_version to be set to 0x50

Type Length Values (TLV) represent individual pieces of data that make up the ADS message conten

Values depicted with a 0x prefix are in Hexidecimal format, meaning that each digit
represents 4 binary bits. Thus, a 2-digit hex number is represented by 1 octet of data. 

The octets are always encoded in Most Significant
Byte (MSB) first order, otherwise known as Big Endian
Encoding. 

![MSB_LSB](https://www.learnpick.in/userfiles/resources_conversion_files/presentation_digital_systems_introduction_1448524485_68825-10.jpg)

References to a NULL setting for a field imply that the field is not carrying a value. However
the NULL octet must still be encoded in the PDU.

The 16-0ctet SMPP Header is a mandatory part of every SMPP PDU and must always be
present. The SMPP PDU Body is optional and may not be included with every SMPP PDU.
There are two ways of viewing a PDU; either as a 16-octet header plus (command_length –
16) octets for the body or as a 4-octet command_length plus (command_length – 4) octets
for the remaining PDU data. 

> SMPP is a binary protocol and also supports asynchronous transmission


## The Operations are described in 6 categories: 
* Session Management 
* Message Submission
* Message Delivery 
* Message Broadcast 
* Anciliary Submission 
* Anciliary Broadcast. 


Bind Operation = MC login request 
