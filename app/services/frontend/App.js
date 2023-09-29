import React from 'react';
import { StatusBar } from 'expo-status-bar';
import { StyleSheet, View, Text } from 'react-native';
import { useFonts } from 'expo-font';

export default function App() {
    const [fontsLoaded] = useFonts({
        'Bai-Jamjuree': require('./assets/fonts/BaiJamjuree-Regular.ttf'),
        'Bai-Jamjuree-Bold': require('./assets/fonts/BaiJamjuree-Bold.ttf'),
    });

    if (!fontsLoaded) {
        return null;
    }

    return (
        <View style={styles.container}>
            <StatusBar style="auto" />
            <Text>BASE</Text>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        fontFamily: 'Bai-Jamjuree',
        flex: 1,
        backgroundColor: '#eee',
        alignItems: 'center',
        justifyContent: 'start',
        paddingTop: 20,
    },
});
