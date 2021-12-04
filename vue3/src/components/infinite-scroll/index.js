import InfiniteScroll from './main'
InfiniteScroll.install = (app) => {
    app.directive('InfiniteScroll', InfiniteScroll)
}
export default InfiniteScroll