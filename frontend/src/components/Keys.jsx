import React from "react"
import { useLoaderData } from "react-router-dom"
import Key from "./Key"

const Keys = () => {
  const { myPublicKey } = useLoaderData()

  return (
    <aside>
      <h2>KEYS</h2>
      <Key publicKey={myPublicKey} name="Mine" me />

      <h4>IMPORTED</h4>
      <Key publicKey={myPublicKey} name="Alice" />
      <Key publicKey={myPublicKey} name="Bob" />

      <div className="import-key">
        <textarea name="newKey" cols="18" rows="4" placeholder="Paste Public Key" />
        <input type="text" name="name" placeholder="Name" />
        <button className="outline">Import Key</button> 
      </div>
    </aside>
  )
}

export default Keys
