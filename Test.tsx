import { VFC } from 'react'

type Props = {
  field1: string
}

export const Test: VFC<Props> = ({ field1 }) => {
  return <div>{field1}</div>
}
