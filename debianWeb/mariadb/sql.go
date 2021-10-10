package mariadb

var (
	QueryUsername = "select username from Ainu.users where username = ?"

	QueryEmail = "select email from Ainu.users where email = ?"

	QueryLoginData = `
select A.id,A.username,A.email,A.password_md5,B.favourite_mode,A.ban_datetime,A.privileges,A.donor_expire,A.status_srting from (select * from Ainu.users where username = ?) A
left join (select id,favourite_mode from Ainu.users_stats where username = ?) B on A.id = B.id`

	QueryUserEmail = "select email from Ainu.users where username = ?"

	QueryAccountVerify = `INSERT INTO DEBIAN.account_verify(id,verify_timeout,key,ip,device_type)VALUES(?,?,?,?,?);`

	QueryLogout = `INSERT INTO DEBIAN.jwt_logout (token_hash,expiration) VALUES (?,?);`

	QueryCheckTokenLogout = `SELECT * FROM DEBIAN.jwt_logout where token_hash = ?;`

	//user_id, verify_key, disable_key
	QueryInsertVerifyCode    = `INSERT INTO DEBIAN.account_verify (user_id, verify_key, disable_key) VALUES (?,?,?) on duplicate key UPDATE verify_key = ?, disable_key = ?`
	QueryVerifyCode          = `SELECT * FROM DEBIAN.account_verify where user_id = ? AND verify_key = ? AND timestamp >= date_sub(now(),INTERVAL 10 MINUTE);`
	QueryDisableCode         = `SELECT * FROM DEBIAN.account_verify where user_id = ? AND disable_key = ?;`
	QueryUpdateAccountStatus = `
update Ainu.users set status_srting = ? where id = ?;
delete from DEBIAN.account_verify where user_id = ?
`

	QueryAPILog = `INSERT INTO DEBIAN.api_log (time, request_id, remote_ip, host, method, uri, user_agent, status, error, latency, latency_human, bytes_in, bytes_out) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?);`

	QueryBeatmapScores = `
select A.id,A.userid,C.username,C.country,A.score,A.rank,A.max_combo,A.mods,A.300_count,A.100_count,A.50_count,A.misses_count,A.time,A.play_mode,A.accuracy,A.pp, B.max_combo
	from (select * from (select * from Ainu.scores where beatmap_id = ? AND play_mode = ? AND score > 0 order by score desc LIMIT 1000000000000000000) A group by userid order by score desc limit ?,?) A
	left Join ( select * from Ainu.beatmaps where beatmap_id = ?) B on A.beatmap_id = B.beatmap_id
	left Join ( select * from Ainu.users_stats ) C on A.userid = C.id
;`
	QueryBeatmapScoresRX = `
select A.id,A.userid,C.username,C.country,A.score,A.rank,A.max_combo,A.mods,A.300_count,A.100_count,A.50_count,A.misses_count,A.time,A.play_mode,A.accuracy,A.pp, B.max_combo
	from (select * from (select * from Ainu.scores_relax where beatmap_id = ? AND play_mode = ? AND pp > 0 order by pp desc LIMIT 1000000000000000000) A group by userid order by pp desc limit ?,?) A
	left Join ( select * from Ainu.beatmaps where beatmap_id = ?) B on A.beatmap_id = B.beatmap_id
	left Join ( select * from Ainu.users_stats ) C on A.userid = C.id
;`

	QueryBeatmapScoresBeatmap = `
select 
	A.beatmap_id,
    A.beatmapset_id,
    A.artist,
    A.title,
    A.version,
    A.%s, -- difficulty_std difficulty_taiko difficulty_ctb difficulty_mania
    A.mode,A.bpm,A.ar,A.cs,A.od,A.hp,
    B.total_length,
    A.hit_length,
    C.playcount,
    C.passcount,
    B.max_combo,
    B.count_circles,
    B.count_spinners,
    B.count_sliders,
    B.creator,
    B.creator_id
    from (select * from Ainu.beatmaps  where beatmap_id = ?) A 
    left join ( select * from BeatmapMirror.beatmaps where id = ? ) B on  A.beatmap_id = B.id,
	(select * from (select count(*) as passcount from Ainu.scores where beatmap_id = ? AND completed = 3) A ,(select count(*) as playcount from Ainu.scores where beatmap_id = ?) B) C
;
`
	QueryBeatmaps = `
select 
	beatmap_id,
    version as Diffname,
    case 
		when mode = 0 then difficulty_std
		when mode = 1 then difficulty_taiko
		when mode = 2 then difficulty_ctb
		when mode = 3 then difficulty_mania
    end as 'Difficulty',
    mode as 'Mode'
 from Ainu.beatmaps  where beatmapset_id = ? ;
`

	)

var (
	QueryRegister01 = `INSERT INTO Ainu.users(username, username_safe, password_md5, salt, email, register_datetime, privileges, password_version, allowedtoreport)VALUES (?, ?,?,'',?,?,?,2,1);`
	QueryRegister02 = `INSERT INTO Ainu.users_stats(id, username, user_color, user_style, ranked_score_std, playcount_std, total_score_std, ranked_score_taiko, playcount_taiko, total_score_taiko, ranked_score_ctb, playcount_ctb, total_score_ctb, ranked_score_mania, playcount_mania, total_score_mania) VALUES (?, ?, 'black', '', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0);`
	QueryRegister03 = `INSERT INTO Ainu.rx_stats (id, username, user_color, user_style, ranked_score_std, playcount_std, total_score_std, ranked_score_taiko, playcount_taiko, total_score_taiko, ranked_score_ctb, playcount_ctb, total_score_ctb, ranked_score_mania, playcount_mania, total_score_mania) VALUES (?, ?, 'black', '', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0);`
	QueryRegister04 = `INSERT INTO Ainu.users_rank (userid) VALUES(?);`
	QueryRegister05 = `INSERT INTO Ainu.rx_rank(userid) VALUES(?);`
)
