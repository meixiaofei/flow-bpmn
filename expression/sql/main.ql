query = fn(query,args...) {
	return sqlctx.Query(query,args...)
}
exec = fn(query,args...) {
	return sqlctx.Exec(query,args...)
}
count = fn(query,args...) {
	return sqlctx.Count(query,args...)
}
one = fn(query,args...) {
	return sqlctx.One(query,args...)
}
querySliceStr = fn(query,key,args...) {
	return SliceStr(sqlctx.Query(query,args...),key)
}
export query,count,one,querySliceStr,exec