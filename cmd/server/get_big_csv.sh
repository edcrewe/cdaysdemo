#!/bin/bash

# Download and create a 1.4 Gb sample CSV file for testing 
# from data.gov.uk CKAN archive of UK Land Registry Price Paid Data
for i in {1995..2002}
do
   wget http://prod.publicdata.landregistry.gov.uk.s3-website-eu-west-1.amazonaws.com/pp-$i.csv -O part_$i.csv
   cat part_*.csv >> big.csv
   rm part_*.csv
done
