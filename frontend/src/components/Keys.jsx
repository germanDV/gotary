import React from "react"
import { useLoaderData } from "react-router-dom"
import Key from "./Key"
import ImportKey from "./ImportKey"

function alphabetically(a, b) {
  return a.name.toLowerCase().localeCompare(b.name.toLowerCase())
}

const Keys = () => {
  const { myPublicKey, contacts } = useLoaderData()
  return (
    <aside>
      <h2>KEYS</h2>
      <Key publicKey={myPublicKey} name="Mine" me />

      <h4>IMPORTED</h4>
      {contacts.sort(alphabetically).map(c => (
        <Key key={c.publicKey.hex} publicKey={c.publicKey.hex} name={c.name} />
      ))}
      
      <ImportKey />
    </aside>
  )
}

export default Keys
