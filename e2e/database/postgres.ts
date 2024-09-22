import {Client} from 'pg'

export const client = new Client({
  user: 'postgres',
  password: 'postgres',
  host: '127.0.0.1',
  port: 5432,
  database: 'app',
})

// const skill = {
//   name: "test"
  
// }

export async function insertData() {
  try {
    await client.connect()
    console.log('Connected to Postgres database')
    // Execute a raw SQL query
    const insertQuery = 'INSERT INTO skill (key, name, description, logo, tags) values ($1, $2, $3, $4, $5)'
    
    client.query(insertQuery,
        [
        "go",
        "test",
        "testDescription",
        "testLogo",
        ["programming language", "system"]
        ])

  } catch (error) {
    console.error('Error connecting to database:', error)
  }
}

export async function deleteData() {
  try {
    await client.query("DELETE FROM skill")
    await client.end()
  } catch (error) {
    console.error('Error connecting to database:', error)
  }
}