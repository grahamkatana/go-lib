import {
    SafeAreaView,
    StyleSheet,
    View,
    Text,
    TextInput,
    ScrollView,
    FlatList,
    Dimensions,
    Image,
    StatusBar,
    TouchableOpacity,
  } from 'react-native';
import React from 'react'
import Foundation from 'react-native-vector-icons/Foundation';
import Loader from '../components/Loader'

const windowWidth = Dimensions.get('window').width;
const windowHeight = Dimensions.get('window').height;

const Scroller = ({item}) => {
  return (
  <>
    <View style={styles.item}>
        <Image source = {{uri:item.image}}
   style = {{ width: '100%', height: '65%',borderRadius:10,borderWidth:0.5, borderColor:'#e9ecef'}}
   />
     <View style={styles.itemBox} >
     <Text numberOfLines={1} style={styles.itemText}>{item.title}</Text>
     </View>
     <Text numberOfLines={1} style={styles.paragraph}>{item.author}</Text>
    </View>
  </>
  )
}


const styles = StyleSheet.create({
    image:{
        height:'100%',
    },
    itemBox:{
        height:25

    },
    paragraph:{
        color:'#252c3a',
        fontWeight:'300',
        padding:4,
        fontSize:12,
        marginLeft:2,
        marginRight:2,
        marginTop:0,

    },
    titleText: {
        fontSize: 24,
        fontWeight: 'bold',
        textAlign: 'center',
        padding: 12
      },
    
      item: {
        backgroundColor: '#fff',
        width: windowWidth*0.32,
        height: 280,
        marginTop:10,
        borderRadius:10,
      },
    
      itemText: {
        fontSize: 13,
        padding:2,
        margin:2,
        fontWeight:'600',
        color: 'black',
      },
  });


export default Scroller