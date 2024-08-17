# blog
**blog** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

## Action

create post
```
blogd tx blog create-post fifa world --from alice --chain-id blog

blogd tx blog create-post fifa world --from bob --chain-id blog
```

view post
```
blogd q blog show-post 0
```

list all post
```
blogd q blog show-post 0
```

list all post by creator
```
blogd q blog list-post-by-creator cosmos1ka27hwmt2pwpwqj9mdvwulf9akqs2zdd4gg0ep
```

update post
```
blogd tx blog update-post "Hello" "blockchain" 0 --from alice --chain-id blog
```

delete post
```
blogd tx blog delete-post 0 --from alice  --chain-id blog
```
