import React from 'react'
import { View,StyleSheet,Text, useWindowDimensions, ActivityIndicator } from 'react-native'

export default function Loader({visible=true}) {
    const {height,width} = useWindowDimensions()
  return (
    visible && <View style={[styles.container,{height,width}]}>
        <View style={styles.loader}>
            <ActivityIndicator size="large" color="#2dba68"/>
            <Text style={styles.text}>Loading...</Text>
        </View>
    </View>
  )
}

const styles = StyleSheet.create({
    container:{
        position:'absolute',
        zIndex:10,
        backgroundColor:'rgba(0,0,0,0.1)',
        justifyContent:'center'
    },
    loader:{
        height:70,
        backgroundColor:'#fff',
        marginHorizontal:50,
        borderRadius:5,
        flexDirection:'row',
        alignItems:'center',
        paddingHorizontal:20,
    },
    text:{
        color:'grey',
        paddingLeft:10,
    }
})
