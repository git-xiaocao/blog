import {Route, Routes, Link, useLocation, useNavigate} from "react-router-dom";
import {Layout, Menu} from 'antd';
import {useEffect} from "react";

import ArticleList from "./Article/ArticleList";
import CategoryList from "./Category/CategoryList";
import UserList from "./User/UserList";

const {Header, Content, Sider} = Layout;

enum PageRoute {
    Article = '/article',
    Category = '/category',
    User = '/user',
}


function Root(): JSX.Element {
    const location = useLocation();
    const navigate = useNavigate();
    let redirect = false;
    switch (location.pathname) {
        case PageRoute.Article:
            break;
        case PageRoute.Category:
            break;
        case PageRoute.User:
            break;
        default:
            redirect = true;
            break;
    }
    useEffect(
        () => {
            //重定向到第一个
            if (redirect) {
                navigate(PageRoute.Article)
            }
        }
    );

    return (
        <Layout>
            <Sider
                style={{height: "100vh"}}
                breakpoint="lg"
                collapsedWidth="0">
                <Menu
                    theme="dark" mode="vertical"
                    defaultSelectedKeys={[redirect ? PageRoute.Article : location.pathname]}
                >
                    <Menu.Item key={PageRoute.Article}>
                        <Link to={PageRoute.Article}>文章</Link>
                    </Menu.Item>
                    <Menu.Item key={PageRoute.Category}>
                        <Link to={PageRoute.Category}>分类</Link>
                    </Menu.Item>
                    <Menu.Item key={PageRoute.User}>
                        <Link to={PageRoute.User}>用户</Link>
                    </Menu.Item>
                </Menu>
            </Sider>
            <Layout>
                <Header className="site-layout-sub-header-background" style={{padding: 0}}/>
                <Content style={{margin: '24px 16px 0'}}>
                    <div style={{padding: 24, minHeight: 360}}>
                        <Routes>
                            <Route path={"/article"} element={<ArticleList/>}/>
                            <Route path={"/category"} element={<CategoryList/>}/>
                            <Route path={"/user"} element={<UserList/>}/>
                            <Route path={"/*"} element={<ArticleList/>}/>
                        </Routes>
                    </div>
                </Content>
            </Layout>
        </Layout>
    );
}

export default Root;