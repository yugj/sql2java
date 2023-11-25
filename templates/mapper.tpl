package {{.BasePackage}}.mapper;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import {{.BasePackage}}.entity.{{.EntityName}}Entity;
import org.springframework.stereotype.Repository;

/**
 * {{.EntityDescription}} mapper
 *
 * @author: {{.Author}}
 * @since: 1.0.0
 */
@Repository
public interface {{.EntityName}}Mapper extends BaseMapper<{{.EntityName}}Entity> {

}