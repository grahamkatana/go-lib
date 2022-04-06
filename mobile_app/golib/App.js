import React from 'react';
import Screens from './src/routes/Screens'
import {
  SafeAreaView,
  StatusBar,
  StyleSheet,
} from 'react-native';


const App = () =>{
return (
 <SafeAreaView style={styles.container}>
   <Screens/>
 </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  container:{
   flex:1,
   width:'100%',
   height:'100%',
   backgroundColor:'#2dba68',
  }
});

export default App;
