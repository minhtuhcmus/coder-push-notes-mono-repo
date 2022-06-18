DELIMITER ;;
CREATE TRIGGER `trig_tbl_notes_before_update`
  BEFORE UPDATE ON `notes` FOR EACH ROW
BEGIN
  IF NEW.user_id != OLD.user_id THEN
    SIGNAL SQLSTATE '45000'
      SET MESSAGE_TEXT = 'tbl_notes.user_id is not allowed to be updated, stop trying to update it.';
END IF;
END;;
DELIMITER ;