import * as React from 'react';
import { Image, Item } from 'semantic-ui-react';

const StoryList: React.SFC<{}> = (props) => {
  return (
    <Item.Group>
      <Item>
        {/*<Item.Image size="tiny" content={<div>test</div>}/>*/}

        <Item.Content>
          <Item.Header as="a">Header</Item.Header>
          <Item.Meta>Description</Item.Meta>
          <Item.Extra>Additional Details</Item.Extra>
        </Item.Content>
      </Item>

      <Item>
        {/*<Item.Image size="tiny" src="/assets/images/wireframe/image.png"/>*/}

        <Item.Content>
          <Item.Header as="a">Header</Item.Header>
          <Item.Meta>Description</Item.Meta>
          <Item.Description>
            <Image src="/assets/images/wireframe/short-paragraph.png"/>
          </Item.Description>
          <Item.Extra>Additional Details</Item.Extra>
        </Item.Content>
      </Item>
    </Item.Group>
  );
};

export default StoryList;
