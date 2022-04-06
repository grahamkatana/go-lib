// In App.js in a new project

import * as React from 'react';
import { View, Text } from 'react-native';

import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import Landing from '../screens/Landing';
import Home  from '../screens/Home';
import PhoneNumber from '../screens/PhoneNumber';


const Stack = createNativeStackNavigator();

function Screens() {
  return (
    <NavigationContainer>
      <Stack.Navigator screenOptions={{headerShown: false,}} initialRouteName={Landing}>
        <Stack.Screen name="Landing" component={Landing} />
        <Stack.Screen name="Home" component={Home} />
        <Stack.Screen name="Phone" component={PhoneNumber} />
      </Stack.Navigator>
    </NavigationContainer>
  );
}

export default Screens;