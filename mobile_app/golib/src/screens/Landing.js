import React,{useState} from 'react'
import {
    SafeAreaView,
    StyleSheet,
    View,
    Text,
    ScrollView,
    Dimensions,
    Image,
    TouchableOpacity,
    StatusBar,
  } from 'react-native';

const windowWidth = Dimensions.get('window').width;
const windowHeight = Dimensions.get('window').height;

const Landing = ({ navigation }) => {

const handleNavigation =()=>{
    navigation.navigate('Home')
}
  return (
    <ScrollView style={styles.container}>
         <StatusBar translucent backgroundColor='#2dba68'/>
        <View style={styles.imagePanel}>
            <Image style={styles.image} source={require('../../assets/onboarding.png')}/>
        </View>
        <View style={styles.textPanel}>
        <Text style={styles.heading}>Go-Lib</Text>
        <Text style={styles.paragraph}>Now you can manage the books
         that you get from the library through our app. Go-Lib makes it easier for you,
         us and all book lovers to have an opportunity to get hands on the books we all need!
         </Text>
         <TouchableOpacity onPressIn={() => {handleNavigation()}} style={styles.button}>
             <Text style={styles.buttonText}>Let's Start</Text>
         </TouchableOpacity>
        </View>
    </ScrollView>
  )
}

const styles = StyleSheet.create({
    container:{
    marginTop:10,
     width:'100%',
     height:'100%',
     backgroundColor:'#fff',
    },
    image:{
        width: '100%',
        height: '100%',
        resizeMode: 'cover',

    },
    buttonText:{
        color:'#fff',
        fontWeight:'600',
        fontSize:18,

    },
    imagePanel:{
        backgroundColor:'#2dba68',
        height:windowHeight*0.70,
        padding:22,
    },
    button:{
        backgroundColor:'#252c3a',
        color:'#fff',
        marginTop:20,
        height:55,
        width:'100%',
        borderRadius:12,
        justifyContent:'center',
        alignItems:'center',

    },
    heading:{
        color:'#1a2232',
        fontWeight:'700',
        fontSize:24,
        marginTop:10,
    },
    paragraph:{
        color:'#252c3a',
        fontWeight:'300',
        padding:2,
        fontSize:15,
        marginTop:10,
    },
    textPanel:{
        height:windowHeight*0.30,
        marginTop:-30,
        backgroundColor:'#fff',
        borderRadius:30,
        padding:22,
    },
  });

export default Landing