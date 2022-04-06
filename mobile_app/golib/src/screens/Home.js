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
  VirtualizedList,
} from 'react-native';
import React, { useEffect, useState } from 'react'
import Foundation from 'react-native-vector-icons/Foundation';
import Loader from '../components/Loader'
import Scroller from '../components/Scroller'
import ScrollerVertical from '../components/ScrollerVertical';
import Seperator from '../components/Seperator'

const windowWidth = Dimensions.get('window').width;
const windowHeight = Dimensions.get('window').height;
const Home = () => {

  const MENUS = [
    {
      id: 1,
      title: 'Popular',
    },
    {
      id: 2,
      title: 'Recent',
    },
    {
      id: 3,
      title: 'Favorites',
    },
  ]



  const BOOKS = [
    {
      id: 1,
      title: 'Clean Code',
      author: 'Robert Cecil Martin',
      image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcToUiz2JxPalZPqoTyAbLMev1L023ckxcV20tz37A5e&usqp=CAE&s',
      quantity: '12',
    },
    {
      id: 2,
      title: 'Head First Java',
      author: 'Bert Bates and Kathy Sierra',
      image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQNSuvEtCh_leEt-3GInxPYD52RQZs4xR4hYAo-cgRMAKPUN-TfJ5Fe03vp&s=5',
      quantity: '45',
    },

    {
      id: 3,
      title: "C Programming Absolute Beginner's Guide",
      author: 'Dean Miller and Greg Perry',
      image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTtaF996ctl_OxYwUwscEGwo_H5v3cT3VyzkVvCGoX1&usqp=CAE&s',
      quantity: '18',
    },

    {
      id: 4,
      title: 'An Introduction to Programming in Go',
      author: 'Caleb Doxsey',
      image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ5izcxkafntYDzm3uhWGgDnZN71JWNbuqWPj3LHPzw&usqp=CAE&s',
      quantity: '3',
    },

    {
      id: 5,
      title: 'React and React Native: Complete Guide to Web and Mobile',
      author: 'Adam Boduch',
      image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQar31qd_hVP7bDj-K6-ah-5oUnKG-yXiuMPeCcEINO&usqp=CAE&s',
      quantity: '16',
    },

    {
      id: 6,
      title: 'Laravel Basics: Creating Web Apps. Its Simple.',
      author: 'Gregory Blake',
      image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcShJljlSQ768myFpDBFFpGXGLATAvSoi9TBVSYo6N3x&usqp=CAE&s',
      quantity: '25',
    },

    {
      id: 7,
      title: 'Learn C#',
      author: 'Chittaranjan Dhurat',
      image: 'https://kbimages1-a.akamaihd.net/7ae0dcb0-f989-4b9c-8a87-5b69b86dccc4/353/569/90/False/c-programming-basics-learn-c-coding-for-beginners-book-1.jpg',
      quantity: '10',
    },
  ]

  return (
    <ScrollView style={styles.container}>
      <StatusBar translucent backgroundColor='#fff' />
      <View style={styles.header}>
        <Foundation name={'align-right'} size={26} color={'#252c3a'} />
        <View style={styles.profileImgContainer}>
          <Image style={styles.image} source={{ uri: 'https://i.pinimg.com/originals/ae/ec/c2/aeecc22a67dac7987a80ac0724658493.jpg' }} />
        </View>
      </View>
      <View style={styles.panel}>
        <Text style={styles.heading}>Hello, User!</Text>
        <Text style={styles.paragraph}>Search your desired books from the library...</Text>
        <View style={styles.searchPanel}>
          <TextInput
            style={styles.search}
            // onChangeText={onChangeNumber}
            // value={number}
            placeholder="Search..."
          />
        </View>
        <Text style={styles.recommendations}>Recommendations</Text>
        <FlatList
          data={BOOKS}
          renderItem={({ item }) => <Scroller item={item} />}
          keyExtractor={item => item.id}
          ItemSeparatorComponent={Seperator}
          horizontal={true}
          showsHorizontalScrollIndicator={false}
        />

        <View style={styles.tabs}>
          {
            MENUS.map((item) => (
              <TouchableOpacity key={item.id} style={styles.button}>
                <Text style={styles.buttonText}>{item.title}</Text>
              </TouchableOpacity>
            ))
          }
        </View>
      </View>
      <SafeAreaView style={styles.allBooks}>
      <FlatList
        data={BOOKS}
        renderItem={({ item }) => <ScrollerVertical item={item} />}
        keyExtractor={item => item.id}
        // ItemSeparatorComponent={Seperator}
        horizontal={false}
        // showsHorizontalScrollIndicator={false}
      />

      </SafeAreaView>
      
    </ScrollView>
  )
}

const styles = StyleSheet.create({
  container: {
    marginTop: 10,
    width: '100%',
    height: '100%',
    backgroundColor: '#fff',
  },
  searchPanel: {
    display: 'flex',
    flexDirection: 'row',

  },
  buttonText: {
    color: '#fff',
    fontSize: 16,
    fontWeight: '500'

  },
  button: {
    borderColor: '#2dba68',
    borderWidth: 1,
    backgroundColor: '#2dba68',
    flex: 1,
    margin: 3,
    height: 45,
    alignItems: 'center',
    justifyContent: 'center',
    borderRadius: 6,
    marginBottom:5

  },
  tabs: {
    flexDirection: 'row',
    justifyContent: 'space-between',

  },
  search: {
    borderWidth: 0.2,
    borderRadius: 5,
    width: '100%',
    padding: 10,
    backgroundColor: '#e9ecef'

  },
  heading: {
    color: 'black',
    fontWeight: '700',
    fontSize: 24,
    marginTop: 10,

  },
  recommendations: {
    color: 'grey',
    fontWeight: '600',
    fontSize: 20,
    marginTop: 15,
    marginBottom: 10,

  },
  paragraph: {
    color: '#252c3a',
    fontWeight: '200',
    padding: 2,
    fontSize: 14,
    marginTop: 10,
    marginBottom: 10,
  },

  profileImgContainer: {
    height: 70,
    backgroundColor: '#fff'
  },

  image: {

    width: 45,
    height: 45,
    borderRadius: 45 * 0.50

  },
  header: {
    backgroundColor: '#fff',
    height: windowHeight * 0.20,
    padding: 22,
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center'

  },
  allBooks:{
    backgroundColor: '#fff',
    borderRadius: 30,
    padding: 22,
    marginTop: -25,

  },
  panel: {
    height: windowHeight * 0.70,
    marginTop: -65,
    backgroundColor: '#fff',
    borderRadius: 30,
    padding: 22,
  },
});

export default Home