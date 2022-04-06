import {
    View,
    Text,
    StyleSheet,
    Dimensions,
    TouchableOpacity,
    Image,
} from 'react-native'
import Fontisto from 'react-native-vector-icons/Fontisto';
import React from 'react'
const windowWidth = Dimensions.get('window').width;
const windowHeight = Dimensions.get('window').height;

const ScrollerVertical = ({ item }) => {
    return (
        <View >
        <TouchableOpacity  style={styles.list}>
        <View style={styles.imageContainer}>
                <Image source={{ uri: item.image }} style={{ width: '100%', height: '100%', borderRadius: 10, borderWidth: 0.5, borderColor: '#e9ecef' }} />
            </View>
            <View style={styles.infoContainer}>
                <View style={styles.info}>
                    <Text numberOfLines={1} style={styles.itemTextPrimary}>{item.title}</Text>
                    <Text numberOfLines={1} style={styles.paragraph}>{item.author}</Text>
                    <Text numberOfLines={1} style={styles.itemText}>{item.quantity} at library</Text>
                </View>
                <Fontisto name='bookmark' color={'grey'} size={18}/>
            </View>
        </TouchableOpacity>
        </View>
    )
}

const styles = StyleSheet.create({

    //
    info:{
        width:'60%',
    },
    infoContainer: {
        width: '75%',
        flexDirection: 'row',
        padding:10,
        justifyContent: 'space-between',


    },
    imageContainer: {
        width: '25%',
        marginRight: 10,

    },
    list: {
        backgroundColor: '#fff',
        flexDirection: 'row',
        height: 100,
        marginBottom: 15,

    },
 
    paragraph: {
        color: '#252c3a',
        fontWeight: '300',
        padding: 4,
        fontSize: 12,
        marginLeft: 2,
        marginRight: 2,
        marginTop: 0,

    },
  
    itemTextPrimary: {
        fontSize: 16,
        padding: 2,
        margin: 2,
        fontWeight: '600',
        color: '#2dba68',
    },

    itemText: {
        fontSize: 13,
        padding: 2,
        margin: 2,
        fontWeight: '600',
        color: 'black',
    },
});



export default ScrollerVertical