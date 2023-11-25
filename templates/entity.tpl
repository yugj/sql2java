package {{.BasePackage}}.entity;

import lombok.Getter;
import lombok.Setter;
import java.math.BigDecimal;
import java.util.Date;
import com.yunx.base.data.mybatis.entity.{{.ParentEntityName}};
import com.baomidou.mybatisplus.annotation.TableName;

/**
 * {{.EntityDescription}} entity
 *
 * @author: {{.Author}}
 * @since: {{.DateTime}}
 */
@Setter
@Getter
@TableName("{{.TableName}}")
public class {{.EntityName}}Entity extends {{.ParentEntityName}}<Long> {

{{ range $index, $field := .EntityFields }}
  /**
   * {{$field.Description}}
   */
  private {{$field.Type}} {{$field.Name}};
 {{ end }}

}