package poc

import (
	"backend/logging"
	"backend/models"
	repo "backend/service"
	"context"
	"database/sql"
	"errors"
)

type pocRepo struct {
	Conn *sql.DB
}

// NewPocRepo retunrs implement of poc dashboard repository interface
func NewPocRepo(Conn *sql.DB) repo.PocRepo {
	return &pocRepo{
		Conn: Conn,
	}
}

// define fetch method
func (m *pocRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Poc, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.Poc, 0)
	for rows.Next() {
		data := new(models.Poc)

		err := rows.Scan(
			&data.ID,
			&data.Account,
			&data.Pocname,
			&data.Technology,
			&data.Objective,
			&data.Owner,
			&data.Teamname,
			&data.Status,
			&data.Remarks,
			&data.Link,
			&data.AssignedTo,
		)
		if err != nil {
			logging.Logger.Errorf(err.Error())
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

// define fetch method
func (m *pocRepo) fetchTecById(ctx context.Context, query string, args ...interface{}) ([]*models.TECTimeline, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.TECTimeline, 0)
	for rows.Next() {
		data := new(models.TECTimeline)

		err := rows.Scan(
			&data.ID,
			&data.TECId,
			&data.Comments,
			&data.UpdatedOn,
		)
		if err != nil {
			logging.Logger.Errorf(err.Error())
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *pocRepo) fetchAccById(ctx context.Context, query string, args ...interface{}) ([]*models.AccTimeline, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.AccTimeline, 0)
	for rows.Next() {
		data := new(models.AccTimeline)

		err := rows.Scan(
			&data.ID,
			&data.AccId,
			&data.Comments,
			&data.UpdatedOn,
		)
		if err != nil {
			logging.Logger.Errorf(err.Error())
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

// define fetchBlogs method
func (m *pocRepo) fetchBlogs(ctx context.Context, query string, args ...interface{}) ([]*models.Blog, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.Blog, 0)
	for rows.Next() {
		data := new(models.Blog)

		err := rows.Scan(
			&data.ID,
			&data.Subject,
			&data.UpdatedOn,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

// define fetchPillar method
func (m *pocRepo) fetchPillar(ctx context.Context, query string, args ...interface{}) ([]*models.Pillars, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.Pillars, 0)
	for rows.Next() {
		data := new(models.Pillars)

		err := rows.Scan(
			&data.ID,
			&data.PillarName,
			&data.CreatedAt,
			&data.IsActive,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

// define fetchTopics method
func (m *pocRepo) fetchTopics(ctx context.Context, query string, args ...interface{}) ([]*models.Topics, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.Topics, 0)
	for rows.Next() {
		data := new(models.Topics)

		err := rows.Scan(
			&data.ID,
			&data.PillarId,
			&data.Topic,
			&data.CreatedAt,
			&data.IsActive,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

// define fetchARBReviews method
func (m *pocRepo) fetchArbReviews(ctx context.Context, query string, args ...interface{}) ([]*models.ARBReviews, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.ARBReviews, 0)
	for rows.Next() {
		data := new(models.ARBReviews)

		err := rows.Scan(
			&data.PillarID,
			&data.PillarName,
			&data.TopicID,
			&data.Topic,
			&data.BestPracticeId,
			&data.BestPractice,
			&data.Description,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}
func (m *pocRepo) fetchSkills(ctx context.Context, query string, args ...interface{}) ([]*models.Skills, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.Skills, 0)
	for rows.Next() {
		data := new(models.Skills)

		err := rows.Scan(
			&data.ID,
			&data.Skill,
			&data.UpdatedOn,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

// define fetchTecMember method
func (m *pocRepo) fetchTecMember(ctx context.Context, query string, args ...interface{}) ([]*models.TecMember, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.TecMember, 0)
	for rows.Next() {
		data := new(models.TecMember)

		err := rows.Scan(
			&data.ID,
			&data.Member,
			&data.Project,
			&data.CoreSkills,
			&data.IsAvailable,
			&data.Comments,
			&data.UpdatedOn,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

// define fetchAccSnap method
func (m *pocRepo) fetchAccSnap(ctx context.Context, query string, args ...interface{}) ([]*models.AccSnap, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.AccSnap, 0)
	for rows.Next() {
		data := new(models.AccSnap)

		err := rows.Scan(
			&data.ID,
			&data.AccName,
			&data.Version,
			&data.IndicativeTimeline,
			&data.ResourceRequirement,
			&data.Blocker,
			&data.Comments,
			&data.UpdatedOn,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

// define fetchAR method
func (m *pocRepo) fetchArb(ctx context.Context, query string, args ...interface{}) ([]*models.Reviews, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.Reviews, 0)
	for rows.Next() {
		data := new(models.Reviews)

		err := rows.Scan(
			&data.ID,
			&data.ProjectName,
			&data.ProjectOwner,
			&data.Reviewer,
			&data.Auditor,
			&data.StatusId,
			&data.StartDate,
			&data.EndDate,
			&data.ProjectScore,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

// define fetchBadge method
func (m *pocRepo) fetchBadge(ctx context.Context, query string, args ...interface{}) ([]*models.Badge, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.Badge, 0)
	for rows.Next() {
		data := new(models.Badge)

		err := rows.Scan(
			&data.ID,
			&data.Badge,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

// define fetchArbStatus method
func (m *pocRepo) fetchArbStatus(ctx context.Context, query string, args ...interface{}) ([]*models.ArbStatus, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.ArbStatus, 0)
	for rows.Next() {
		data := new(models.ArbStatus)

		err := rows.Scan(
			&data.ID,
			&data.Status,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

// Get Pocs list
func (m *pocRepo) FetchPocs(ctx context.Context) ([]*models.Poc, error) {
	query := `Select id, account, pocname, IFNULL(technology, '') as technology, 
	objective, owner, teamname, status, remarks, IFNULL(link, '') as link, 
	IFNULL(assignedto, '') as assignedto From pocs WHERE isactive=true ORDER BY id desc;`
	return m.fetch(ctx, query)
}

// Get feed list
func (m *pocRepo) FetchBlogs(ctx context.Context) ([]*models.Blog, error) {
	query := `Select id, subject, updatedon From blog ORDER BY id desc;`
	return m.fetchBlogs(ctx, query)
}

// Get pillar list
func (m *pocRepo) FetchPillars(ctx context.Context) ([]*models.Pillars, error) {
	query := `Select * From pillarmaster WHERE isactive=true;`
	return m.fetchPillar(ctx, query)
}

// Get pillar list
func (m *pocRepo) FetchTopics(ctx context.Context) ([]*models.Topics, error) {
	query := `SELECT ROW_NUMBER() OVER (ORDER BY id) AS sno, id, pillarid, topic, isactive, createdat From topicmaster WHERE isactive=true;`
	return m.fetchTopics(ctx, query)
}

// Get ARB Reviews list
func (m *pocRepo) FetchArbReviews(ctx context.Context) ([]*models.ARBReviews, error) {
	query := `select p.id as pillarid, COALESCE(p.pillarname, '') AS pillarname, t.id as topicid,  COALESCE(t.topic, '') AS topic,
	COALESCE(q.id, -1) AS bestpracticeid, COALESCE(q.bestpractice, '') AS bestpractice, COALESCE(q.description, '') as description from pillarmaster p left join topicmaster t on p.id = t.pillarid
	left join bestpracticemaster q on t.id= q.topicid order by p.id asc, topicid asc, q.id asc;`
	return m.fetchArbReviews(ctx, query)
}

// Get TEC member list
func (m *pocRepo) FetchTecMembers(ctx context.Context) ([]*models.TecMember, error) {
	query := `select id, member, IFNULL(project, '') as project, IFNULL(coreskills, '') as coreskills, isavailable, IFNULL(comments, '') as comments, updatedon FROM tecmember order by id desc;`
	return m.fetchTecMember(ctx, query)
}

// Get TEC member list
func (m *pocRepo) FetchSkills(ctx context.Context) ([]*models.Skills, error) {
	query := `select id, skill, updatedon from skillmaster where isactive=1 ORDER by ID desc;`
	return m.fetchSkills(ctx, query)
}

// Get TEC member list
func (m *pocRepo) FetchTecActivity(ctx context.Context) ([]*models.TecMember, error) {
	query := `select id, member, IFNULL(project, '') as project, isavailable, IFNULL(comments, '') as comments, updatedon FROM tecmember order by id desc;`
	return m.fetchTecMember(ctx, query)
}

// Get Accelerator Snapshot list
func (m *pocRepo) FetchAccSnap(ctx context.Context) ([]*models.AccSnap, error) {
	query := `SELECT id, accname, IFNULL(version, '') as version, IFNULL(indicativetimeline, '') as indicativetimeline,
		IFNULL(resourcerequirement, '') as resourcerequirement, IFNULL(blocker, '') as blocker, IFNULL(comments, '') as comments,
			updatedon FROM acceleratorsnap order by id desc;`
	return m.fetchAccSnap(ctx, query)
}

// Get Arb list
func (m *pocRepo) FetchArb(ctx context.Context) ([]*models.Reviews, error) {
	query := `SELECT a.id, projectname, projectowner, reviewer, auditor, statusid, IFNULL(startdate, '') as startdate,
		IFNULL(enddate, '') as enddate, projectscore From reviewboard a LEFT JOIN arbstatusmaster b on a.statusid = b.id
		 WHERE a.isactive=true ORDER BY a.id desc;`
	return m.fetchArb(ctx, query)
}

// Get badge list
func (m *pocRepo) FetchBadges(ctx context.Context) ([]*models.Badge, error) {
	query := `SELECT id, badge FROM badgemaster WHERE isactive=true ORDER BY id desc;`
	return m.fetchBadge(ctx, query)
}

// Get arb status list
func (m *pocRepo) FetchArbStatus(ctx context.Context) ([]*models.ArbStatus, error) {
	query := `SELECT id, status FROM arbstatusmaster WHERE isactive=true;`
	return m.fetchArbStatus(ctx, query)
}

// Get Poc by id
func (m *pocRepo) GetPocByID(ctx context.Context, id int64) (*models.Poc, error) {
	query := `Select id, account, pocname FROM pocs WHERE id=?`

	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}

	payload := &models.Poc{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return payload, errors.New("requested poc is not found")
	}

	return payload, nil
}

// Get Poc by id
func (m *pocRepo) GetTecActivityByID(ctx context.Context, id int64) ([]*models.TECTimeline, error) {
	query := `Select id, tecid, comments, updatedon FROM tectimeline WHERE tecid=? order by id desc;`
	return m.fetchTecById(ctx, query, id)
}

// Get Poc by id
func (m *pocRepo) GetAccActivityByID(ctx context.Context, id int64) ([]*models.AccTimeline, error) {
	query := `Select id, accid, comments, updatedon FROM acctimeline WHERE accid=? order by id desc;`
	return m.fetchAccById(ctx, query, id)
}

// Create new Poc
func (m *pocRepo) CreatePoc(ctx context.Context, p *models.Poc) (int64, error) {
	query := `INSERT INTO pocs (account, pocname, technology, objective, owner, teamname, status, remarks,
								link, assignedto) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ? )`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return -1, err
	}
	res, err := stmt.ExecContext(ctx, p.Account, p.Pocname, p.Technology, p.Objective, p.Owner, p.Teamname, p.Status, p.Remarks, p.Link, p.AssignedTo)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return -1, err
	}

	return res.RowsAffected()
}

// Create new TecMember
func (m *pocRepo) CreateTecMember(ctx context.Context, p *models.SaveTecMember) (int64, error) {
	query := `INSERT INTO tecmember (member, project, coreskills, comments, isavailable, updatedon) VALUES(?, ?, ?, ?, ?, ?)`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return -1, err
	}
	res, err := stmt.ExecContext(ctx, p.Member, p.Project, p.CoreSkills, p.Comments, p.IsAvailable, p.UpdatedOn)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return -1, err
	}

	return res.RowsAffected()
}

// Create new TecMember
func (m *pocRepo) CreateBlog(ctx context.Context, p *models.Blog) (int64, error) {
	query := `INSERT INTO blog (subject, updatedon) VALUES(?, ?)`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return -1, err
	}
	res, err := stmt.ExecContext(ctx, p.Subject, p.UpdatedOn)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return -1, err
	}

	return res.RowsAffected()
}

// Create new Core skill
func (m *pocRepo) CreateCoreSkill(ctx context.Context, p *models.Skills) (int64, error) {
	query := `INSERT INTO skillmaster (skill, updatedon) VALUES(?, now())`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return -1, err
	}
	res, err := stmt.ExecContext(ctx, p.Skill)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return -1, err
	}
	return res.RowsAffected()
}

// Update core skill
func (m *pocRepo) UpdateCoreSkill(ctx context.Context, p *models.Skills) (*models.Skills, error) {
	query := "UPDATE skillmaster set skill=?, updatedon=now() where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Skill,
		p.ID,
	)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

// Create new TecMember
func (m *pocRepo) CreateTecTimeline(ctx context.Context, p *models.TECTimeline) (int64, error) {
	query := `INSERT INTO tectimeline (tecid, comments, updatedon) VALUES(?, ?, ?)`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return -1, err
	}
	res, err := stmt.ExecContext(ctx, p.TECId, p.Comments, p.UpdatedOn)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return -1, err
	}

	return res.RowsAffected()
}

func (m *pocRepo) CreateAccTimeline(ctx context.Context, p *models.AccTimeline) (int64, error) {
	query := `INSERT INTO acctimeline (accid, comments, updatedon) VALUES(?, ?, ?)`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return -1, err
	}
	res, err := stmt.ExecContext(ctx, p.AccId, p.Comments, p.UpdatedOn)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return -1, err
	}

	return res.RowsAffected()
}

// Create new Acc snap
func (m *pocRepo) CreateAccSnap(ctx context.Context, p *models.AccSnap) (int64, error) {
	query := `INSERT INTO acceleratorsnap (accname, version, indicativetimeline, resourcerequirement, blocker, comments, updatedon) VALUES(?, ?, ?, ?, ?, ?, ?)`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return -1, err
	}
	res, err := stmt.ExecContext(ctx, p.AccName, p.Version, p.IndicativeTimeline, p.ResourceRequirement, p.Blocker, p.Comments, p.UpdatedOn)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return -1, err
	}

	return res.RowsAffected()
}

// Update TecMember
func (m *pocRepo) UpdateBlog(ctx context.Context, p *models.Blog) (*models.Blog, error) {
	query := "UPDATE blog set subject=?, updatedon=? where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Subject,
		p.UpdatedOn,
		p.ID,
	)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

// Update TecMember
func (m *pocRepo) UpdateTecMember(ctx context.Context, p *models.SaveTecMember) (*models.SaveTecMember, error) {
	query := "UPDATE tecmember set member=?, project=?, coreskills=?, comments=?, isavailable=?, updatedon=? where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Member,
		p.Project,
		p.CoreSkills,
		p.Comments,
		p.IsAvailable,
		p.UpdatedOn,
		p.ID,
	)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

// Update Accsnap
func (m *pocRepo) UpdateAccSnap(ctx context.Context, p *models.AccSnap) (*models.AccSnap, error) {
	query := "UPDATE acceleratorsnap set accname=?, version=?, indicativetimeline=?, resourcerequirement=?, blocker=?,comments=?, updatedon=? where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.AccName,
		p.Version,
		p.IndicativeTimeline,
		p.ResourceRequirement,
		p.Blocker,
		p.Comments,
		p.UpdatedOn,
		p.ID,
	)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

// Update Poc
func (m *pocRepo) UpdatePoc(ctx context.Context, p *models.Poc, id int64) (*models.Poc, error) {
	query := "UPDATE pocs set account=?, pocname=?, technology=?, objective=?, owner=?, teamname=?, status=?, remarks=?, link=?, assignedto=? where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Account,
		p.Pocname,
		p.Technology,
		p.Objective,
		p.Owner,
		p.Teamname,
		p.Status,
		p.Remarks,
		p.Link,
		p.AssignedTo,
		id,
	)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

// Update Arb
func (m *pocRepo) UpdateArb(ctx context.Context, p *models.Reviews, id int64) (*models.Reviews, error) {
	query := "UPDATE reviewboard set projectname=?, projectowner=?, reviewer=?, auditor=?, startdate=?, enddate=?, projectscore=? where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.ProjectName,
		p.ProjectOwner,
		p.Reviewer,
		p.Auditor,
		p.StartDate,
		p.EndDate,
		p.ProjectScore,
		id,
	)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

// Delete Poc
func (m *pocRepo) DeletePoc(ctx context.Context, id int64) (int64, error) {
	query := "UPDATE pocs SET isactive = 0 WHERE id=?"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return 0, err
	}
	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return 0, err
	}
	return res.RowsAffected()
}

// Delete Poc
func (m *pocRepo) DeleteFeed(ctx context.Context, id int64) (int64, error) {
	query := "UPDATE feed SET isactive = 0 WHERE id=?"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return 0, err
	}
	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return 0, err
	}
	return res.RowsAffected()
}

// define fetch method
func (m *pocRepo) fetchTechCount(ctx context.Context, query string, args ...interface{}) ([]*models.TechCount, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.TechCount, 0)
	for rows.Next() {
		data := new(models.TechCount)

		err := rows.Scan(
			&data.Technology,
			&data.Count,
		)
		if err != nil {
			logging.Logger.Errorf(err.Error())
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

// Get techwise list
func (m *pocRepo) FetchTechCount(ctx context.Context) ([]*models.TechCount, error) {
	query := `SELECT technology, count(*) as count FROM pocs WHERE isactive=1 GROUP BY technology;`
	return m.fetchTechCount(ctx, query)
}

// Get piechartcount list
func (m *pocRepo) FetchPieChartCount(ctx context.Context) ([]*models.PieChartCount, error) {
	query := `SELECT account, count(*) as count FROM pocs WHERE isactive=1 GROUP BY account;`
	return m.fetchPieChartCount(ctx, query)
}

// define fetchPieChartCount method
func (m *pocRepo) fetchPieChartCount(ctx context.Context, query string, args ...interface{}) ([]*models.PieChartCount, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.PieChartCount, 0)
	for rows.Next() {
		data := new(models.PieChartCount)

		err := rows.Scan(
			&data.Account,
			&data.Count,
		)
		if err != nil {
			logging.Logger.Errorf(err.Error())
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

// Update Feed
func (m *pocRepo) UpdateArbStatus(ctx context.Context, statusid uint, id int64) (int64, error) {
	query := "UPDATE reviewboard SET statusid=? WHERE id=?"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return 0, err
	}
	res, err := stmt.ExecContext(ctx, statusid, id)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return 0, err
	}
	return res.RowsAffected()
}

// Update Feed
func (m *pocRepo) UpdateDescription(ctx context.Context, description string, id int64) (int64, error) {
	query := "UPDATE bestpracticemaster SET description=? WHERE id=?"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return 0, err
	}
	res, err := stmt.ExecContext(ctx, description, id)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return 0, err
	}
	return res.RowsAffected()
}

// Create new Poc
func (m *pocRepo) CreateArb(ctx context.Context, p *models.Reviews) (int64, error) {
	query := `INSERT INTO reviewboard (projectname, projectowner, reviewer, auditor, projectscore, startdate, enddate)
	 VALUES(?, ?, ?, ?, ?, ?, ?)`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return -1, err
	}
	res, err := stmt.ExecContext(ctx, p.ProjectName, p.ProjectOwner, p.Reviewer, p.Auditor, p.ProjectScore, p.StartDate, p.EndDate)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return -1, err
	}

	return res.RowsAffected()
}

// Delete Poc
func (m *pocRepo) DeleteArb(ctx context.Context, id int64) (int64, error) {
	query := "UPDATE reviewboard SET isactive=0 WHERE id=?"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return 0, err
	}
	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		return 0, err
	}
	return res.RowsAffected()
}
