
set -e

DB_USER='root';
DB_PASSWORD='1234';
DB_HOST='localhost';
DB_PORT=3306;
DB_NAME='rickandmorty';


TOTAL_PAGES=$(curl -s https://rickandmortyapi.com/api/character | jq -r .info.pages)

i=1

while [ "$i" -le ${TOTAL_PAGES} ]; do
                
       curl --request GET -s "https://rickandmortyapi.com/api/character?page=${2}" | jq -r '.results' > res.json
        
        jq -c '.[]' res.json | while read j; do 
                
               
                

            
                idcharacter=$(echo ${j}  | jq -r '.id')
                echo $idcharacter
                nombre=$(echo ${j}  | jq -r '.name')
                echo $nombre
                estado=$(echo ${j}  | jq -r '.status')
                echo $estado
                especie=$(echo ${j}  | jq -r '.species')
                echo $especie
               
                
                echo "INSERT INTO rickandmorty.character (idcharacter,nombre,estado,especie) VALUES ('${idcharacter//\'/\'}','${nombre//\'/\'}','${estado//\'/\'}','${especie}');"  | mysql --user=$DB_USER --password=$DB_PASSWORD --host=$DB_HOST --port=$DB_PORT $DB_NAME
             
             
              
                 
        done 
        i=$((i+1));
        break   
        
done

TOTAL_PAGES2=$(curl -s https://rickandmortyapi.com/api/location | jq -r .info.pages)

l=1

while [ "$l" -le ${TOTAL_PAGES2} ]; do
                
       curl --request GET -s "https://rickandmortyapi.com/api/location?page=${2}" | jq -r '.results' > res.json
        
        jq -c '.[]' res.json | while read k; do 
                
            
                idlocation=$(echo ${k}  | jq -r '.id')
                echo $idlocation
                nombre=$(echo ${k}  | jq -r '.name')
                nombrerp=$(echo "$nombre" | sed -e   's/[[:punct:]]/'$replace'/g')
                echo $nombrerp
  
                #echo $nombre
                tipo=$(echo ${k}  | jq -r '.type')
                echo $tipo
                dimension=$(echo ${k}  | jq -r '.dimension')
                echo $dimension
               
                
                echo "INSERT INTO rickandmorty.location(idlocation,nombre,tipo,dimension) VALUES ('${idlocation//\'/\'}','${nombrerp//\'/\'}','${tipo//\'/\'}','${dimension}');"  | mysql --user=$DB_USER --password=$DB_PASSWORD --host=$DB_HOST --port=$DB_PORT $DB_NAME
                
                 
        done 
        l=$((l+1));
        break    
done


exec "$@"

