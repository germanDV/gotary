import React from "react"
import { useFetcher } from "react-router-dom"
import toast from "react-simple-toasts"
import ClipboardIcon from "../icons/ClipboardIcon"
import TrashIcon from "../icons/TrashIcon"
import { CopyToClipboard } from "../../wailsjs/go/main/App"

const Key = ({ publicKey, name, me }) => {
  const fetcher = useFetcher()

  const copyToClipboard = async () => {
    try {
      await CopyToClipboard(publicKey)
      toast("Copied to clipboard!")
    } catch {
      toast("Error trying to copy key")
    }
  }

  const handleDelete = () => {
    fetcher.submit({ name }, { method: "delete" })
  }

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
          <div className="danger" title="Delete" onClick={handleDelete}>
            <TrashIcon />
          </div>
        ) : null}
      </div>
    </div>
  ) 
}

export default Key
