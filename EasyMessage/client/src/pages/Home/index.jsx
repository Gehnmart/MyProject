import React from 'react';
import { Message } from 'components'


const Home = () => {
    return (
        <div className='home'>
            <Message
                avatar={"https://sun6-21.userapi.com/s/v1/if2/DzKTqHRoFr-77-58W7ZIZ7i98nYSVViSjy5ekV_aWJqgk-KJUy38ujDNNB_8J7ekRp30N28MZNqYtiMwBe0ntskj.jpg?quality=95&crop=131,140,720,720&as=50x50,100x100,200x200,400x400&ava=1&u=ubOL4DVTvUTtucxR5YN7MgdrKsKXZEboR40I1NPKN44&cs=200x200"}
                text='Привеt ,mas dmnf,mans чел!'
                date='Sun Feb 24 2022, 17:58:00'
                user={{ name: 'Вася Пупкин', id: 1 }}>
            </Message>
            <Message
                avatar={"https://sun6-21.userapi.com/s/v1/if2/DzKTqHRoFr-77-58W7ZIZ7i98nYSVViSjy5ekV_aWJqgk-KJUy38ujDNNB_8J7ekRp30N28MZNqYtiMwBe0ntskj.jpg?quality=95&crop=131,140,720,720&as=50x50,100x100,200x200,400x400&ava=1&u=ubOL4DVTvUTtucxR5YN7MgdrKsKXZEboR40I1NPKN44&cs=200x200"}
                text='Привеt ,mas dmnf,mans чел!'
                date='Sun Feb 24 2022, 17:58:00'
                user={{ name: 'Вася Пупкин', id: 1 }}
                is_me={true}>
            </Message>
            <Message
                avatar={"https://sun6-21.userapi.com/s/v1/if2/DzKTqHRoFr-77-58W7ZIZ7i98nYSVViSjy5ekV_aWJqgk-KJUy38ujDNNB_8J7ekRp30N28MZNqYtiMwBe0ntskj.jpg?quality=95&crop=131,140,720,720&as=50x50,100x100,200x200,400x400&ava=1&u=ubOL4DVTvUTtucxR5YN7MgdrKsKXZEboR40I1NPKN44&cs=200x200"}
                text='Привеt ,mas dmnf,mans чел!'
                date='Sun Feb 24 2022, 17:58:00'
                user={{ name: 'Вася Пупкин', id: 1 }}>
            </Message>
        </div>
    );
}

export default Home;
