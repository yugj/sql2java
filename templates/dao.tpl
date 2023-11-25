package {{.BasePackage}}.dao;

import com.yunx.base.data.mybatis.repository.AbstractDao;
import {{.BasePackage}}.entity.{{.EntityName}}Entity;
import {{.BasePackage}}.mapper.{{.EntityName}}Mapper;
import org.springframework.stereotype.Service;

/**
 * {{.EntityDescription}} dao
 *
 * @author: {{.Author}}
 * @since: 1.0.0
 */
@Service
public class {{.EntityName}}Dao extends AbstractDao<{{.EntityName}}Mapper, {{.EntityName}}Entity> {

}