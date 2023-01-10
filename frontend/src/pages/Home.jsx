import React from "react"
import { Outlet } from "react-router-dom"
import Keys from "../components/Keys"
import Nav from "../components/Nav"
import { GetMyPublicKey, GetContacts, AddContact, DeleteContact } from "../../wailsjs/go/main/App"

const Home = () => {
  return (
    <main>
      <div className="grid">
        <Keys />
        <div className="main">
          <Nav />
          <Outlet />
        </div>
      </div>
    </main>
  )
}

export default Home

export async function loader() {
  const [myPublicKey, contacts] = await Promise.all([GetMyPublicKey(), GetContacts()]) 
  return { myPublicKey, contacts }
}

export async function action({ request }) {
  if (request.method === "POST") {
    const formData = await request.formData()
    const name = formData.get("newName")
    const key = formData.get("newKey")
    return AddContact(name, key)
  } else if (request.method === "DELETE") {
    const formData = await request.formData()
    const name = formData.get("name")
    return DeleteContact(name)
  }
}
