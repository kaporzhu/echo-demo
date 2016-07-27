package demo

// 没有把LoginRequired作为中间件，是因为LoginRequired里面涉及了具体业务层的东西，比如用户信息。
// middleware提供的应该是独立于具体项目的中间件。
