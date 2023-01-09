import React from "react"
import Keys from "../components/Keys"
import Signer from "../components/Signer"
import Out from "../components/Out"
import { GetMyPublicKey, GetContacts, AddContact } from "../../wailsjs/go/main/App"

const Home = () => {
  return (
    <main>
      <Out />
      <div className="grid">
        <Keys />
        <div className="main">
          <Signer />
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

export async function importKey({ request }) {
  const formData = await request.formData()
  const name = formData.get("newName")
  const key = formData.get("newKey")
  return AddContact(name, key)
}
