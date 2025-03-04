#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
WHITE='\033[0;37m'
YELLOW='\033[0;33m'
CYAN='\033[1;36m'
NC='\033[0m'

echo -e "${GREEN}Ishvalan region detected.${NC}"
sleep 2 # Kasih delay biar bagus

# Bikin persentase progress
printf "${YELLOW}Scanning...${NC} [  0%% ]"
for i in $(seq 1 100); do
    printf "\r${YELLOW}Scanning...${NC} [ %3d%% ]" "$i"
    sleep 0.02
done
printf "\n${WHITE}100.000${NC} ${GREEN}Ishvalan detected.${NC}\n"
sleep 2

# King Bradley nyuruh
echo -e "${RED}King Bradley has launched the War of Extermination on Ishvalan!${NC}"
sleep 2

# Bawahannya udah dimasukkin /etc/sudoers
echo -e "${GREEN}Superuser privilege acquired.${NC}"
sleep 1

# Persentase proggress tapi ada stutteringnya, biar kayak ngehapus file beneran
printf "${YELLOW}Executing:${NC} ${CYAN}sudo rm -rf Ishvalan${NC} [  0%% ]"
progress=0
while [ "$progress" -lt 100 ]; do
    delay=$(awk -v min=0.1 -v max=0.5 'BEGIN{srand(); print min+rand()*(max-min)}')
    jump=$(awk -v min=1 -v max=5 'BEGIN{srand(); print int(min+rand()*(max-min+1))}')

    progress=$((progress + jump))
    if [ "$progress" -gt 100 ]; then progress=100; fi

    printf "\r${YELLOW}Executing:${NC} ${CYAN}sudo rm -rf Ishvalan${NC} [ %3d%% ]" "$progress"
    sleep "$delay"
done
printf "\n"
sleep 2

echo -e "${YELLOW}Processing human sacrifices...${NC}"
sleep 3
echo -e "${YELLOW}Creating Philosopher's Stone...${NC}"
sleep 2

echo -e "${GREEN}Philosopher's Stone has been created.${NC}"

echo -e "${RED}"
cat << "EOF"
      .     .
   .  |     |  .
  | ~ .:::::. ~ |
  |  :::::::::  |
  |  :::::::::  |
  |  ':::::::'  |
  |    ':::'    |
   .     '     .
    ' - ._. - '
EOF
echo -e "${NC}"