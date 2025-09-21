# CRUD Operation Go with Postgres

#### Create and Run Go
1. Install Go on windows
	- `go version`
2. Install VS code
3. Install Extensions
	- golang.go
	- formulahendry.code-runner
	- esbenpp.pretier-vscodee
	- yzhang.markdown-all-in-one
4. Create main.go file
5. Install all the packages
	- `go mod init {module-name}`
	- `go run main.go`

#### Setup Docker Postgres with Network to connect with PgAdmin
`docker network create pgnetwork`
`docker run --name my-postgres --network pgnetwork -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=password -e POSTGRES_DB=test_db -p 5432:5432 -d postgres`
`docker exec -it my-postgres psql -U admin -d test_db`
- Run `\dt` → see tables.    
- Run `\d books` → see your `books` table structure.
- Run `SELECT * FROM public.books;` → query data.

#### Optional Installm pgAdmin to see postgres db 
`docker run --name my-pgadmin --network pgnetwork -e PGADMIN_DEFAULT_EMAIL=admin@admin.com -e PGADMIN_DEFAULT_PASSWORD=admin -p 5050:80 -d dpage/pgadmin4`

#### If any process is listening in the same port kill that
- `netstat -ano | findstr :3000`  
	- ex output: TCP    0.0.0.0:3000    0.0.0.0:0    LISTENING    12345  
- `taskkill /PID 12345 /F`

#### Previewing Your Markdown File
Once the extension is installed:
1. Open your `.md` file in VS Code.
2. To open the preview, press `Ctrl+Shift+V` (Windows/Linux) or `Cmd+Shift+V` (macOS).
3. To toggle between the editor and preview, press `Ctrl+K V` (Windows/Linux) or `Cmd+K V` (macOS).


#### Git for Source Control
1. Install Git (https://git-scm.com/downloads/win)
    - `git --version`
2. `git config --global user.name {Your Name}`
3. `git config --global user.email {your.email@example.com}`
4. Restart VS code
5. Create a repository in your git
6. `git remote add origin {URL of your git repository}`
7. `git commit -m "first commit"`
8. `git push -u origin main`