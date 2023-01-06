import React from "react"
import toast from "react-simple-toasts"
import ClipboardIcon from "../icons/ClipboardIcon"
import TrashIcon from "../icons/TrashIcon"
import { CopyToClipboard } from "../../wailsjs/go/main/App"

const Key = ({ publicKey, name, me }) => {
  const copyToClipboard = async () => {
    try {
      await CopyToClipboard(publicKey)
      toast("Copied to clipboard!")
    } catch {
      toast("Error trying to copy key")
    }
  }

  const remove = () => alert(publicKey)

  return (
    <div className="key">
      <div className="name">
        <span>{name}</span>
        <span className="public-key-xs">
          {publicKey.substring(0, 8)}
          ...
          {publicKey.substring(56)}
        </span>
      </div>
      <div className="actions">
        <div onClick={copyToClipboard} title="Copy">
          <ClipboardIcon />
        </div>
        {!me ? (
          <div className="danger" onClick={remove} title="Delete">
          <TrashIcon />
          </div>
        ) : null}
      </div>
    </div>
  ) 
}

export default Key
