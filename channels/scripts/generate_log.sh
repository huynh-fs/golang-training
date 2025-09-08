#!/bin/bash

OUTPUT_FILE="large.log"
NUM_LINES=100000

echo "Generating a large log file with $NUM_LINES lines..."

rm -f $OUTPUT_FILE

declare -a levels=("INFO" "WARN" "ERROR" "DEBUG")
declare -a messages=(
  "User logged in successfully"
  "Database connection established"
  "Cache miss for key: user:123"
  "Failed to process payment"
  "Request timeout"
  "Invalid credentials provided"
  "Starting background job"
  "Configuration loaded"
)

for (( i=1; i<=$NUM_LINES; i++ ))
do
  RANDOM_LEVEL=${levels[$RANDOM % ${#levels[@]}]}
  RANDOM_MESSAGE=${messages[$RANDOM % ${#messages[@]}]}
  
  UUID=$(date +%s%N)$RANDOM
  echo "$RANDOM_LEVEL: $RANDOM_MESSAGE - transaction_id=$UUID" >> $OUTPUT_FILE
done

echo "Done. Log file created: $OUTPUT_FILE"