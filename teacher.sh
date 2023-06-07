#! /bin/bash
numInterview=$(head -n 179 streets/Buckingham_Place|tail -n 1|cut -d "#" -f2)
echo "$numInterview"
cat ./interviews/interview-"$numInterview"
echo "$MAIN_SUSPECT"
