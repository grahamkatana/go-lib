import { View, Text } from 'react-native'
import React from 'react'

const Seperator = ({height=1,width=10}) => {
  return (
    <View
    style={{
      height: height,
      width: width,
      backgroundColor: "white",
    }}
  />
  )
}

export default Seperator