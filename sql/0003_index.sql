USE blog;

-- Comments are looked up per post_id
ALTER TABLE comments ADD INDEX(post_id);

-- TiFlash ColumnStore
/*T! ALTER TABLE posts SET TIFLASH REPLICA 1 */;
/*T! ALTER TABLE comments SET TIFLASH REPLICA 1 */;
