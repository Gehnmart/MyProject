import React from 'react';
import { Button, Message, Dialog } from 'components';
import {TeamOutlined, FormOutlined, SendOutlined} from '@ant-design/icons';
import {Input} from 'antd';

import './Home.scss';
import name_to_rgb from "../../helpers/name_to_rgb";

const { Search } = Input;

const Home = () => {
    const [currentDialog, setCurrentDialog] = React.useState('');
    const [search, setSearch] = React.useState('');
    // Функция для фильтрации диалогов по имени пользователя
    const handleSearch = (items) => {
        if (search) {
            return items.filter(item => item.user.name.toLowerCase().includes(search.toLowerCase()));
        }
        return items;
    };

    const sendButton = () => {
        return <Button type="primary" style={{width: '40px', height: '40px', lineHeight: '2'}} icon={<SendOutlined />} />;
    };

    return (
        <section className='home'>
            <div className="chat">
                <div className="chat__sidebar">
                    <div className="chat__sidebar-header">
                        <div>
                            <TeamOutlined />
                            <span>Список диалогов</span>
                        </div>
                        <FormOutlined />
                    </div>
                    <div className="chat__sidebar-search">
                        <Input value={search} onChange={e => setSearch(e.target.value)} />
                    </div>
                    <div className="chat__sidebar-dialogs">
                        {/* Передаем отфильтрованные диалоги в компонент Dialog */}
                        <Dialog items={handleSearch( [
                            { user: { name: 'Алексей Кривченко', id: 1 }, message: 'Привет, как дела?' },
                            { user: { name: 'Иван Петрович', id: 1 }, message: 'У меня проблемы по сайту, прошу помочь' },
                            { user: { name: 'Аят Бугаевич', id: 1 }, message: 'Что делаешь там? Надо помочь' }
                        ])} />
                    </div>
                </div>
                <div className="chat__dialog">
                    <div className="chat__dialog-header">
                        <b className='chat__dialog-header-name'>
                            {currentDialog}
                        </b>
                    </div>
                    <div className="chat__dialog-messages">
                        <Message
                            is_me={true}
                            user={{ name: 'Алексей', id: 1 }}
                            avatar='https://sun6-21.userapi.com/s/v1/if2/DzKTqHRoFr-77-58W7ZIZ7i98nYSVViSjy5ekV_aWJqgk-KJUy38ujDNNB_8J7ekRp30N28MZNqYtiMwBe0ntskj.jpg?quality=95&crop=131,140,720,720&as=50x50,100x100,200x200,400x400&ava=1&u=ubOL4DVTvUTtucxR5YN7MgdrKsKXZEboR40I1NPKN44&cs=200x200'
                            date="Jan 1 2022 00:00"
                            text="Привет, как дела?"
                        />
                        <Message
                            is_me={false}
                            user={{ name: 'Алексей', id: 1 }}
                            avatar='https://sun6-21.userapi.com/s/v1/if2/DzKTqHRoFr-77-58W7ZIZ7i98nYSVViSjy5ekV_aWJqgk-KJUy38ujDNNB_8J7ekRp30N28MZNqYtiMwBe0ntskj.jpg?quality=95&crop=131,140,720,720&as=50x50,100x100,200x200,400x400&ava=1&u=ubOL4DVTvUTtucxR5YN7MgdrKsKXZEboR40I1NPKN44&cs=200x200'
                            date="Jan 1 2022 00:00"
                            text="Привет, как дела как семья как жизнь как мать как дети?"
                        />
                    </div>
                    <div className="chat__dialog-input">
                        <Input placeholder="Введите сообщение" suffix={sendButton()} />
                    </div>
                </div>
            </div>
        </section>
    );
};

export default Home;
