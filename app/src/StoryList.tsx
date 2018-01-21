import * as React from 'react';
import { Item } from 'semantic-ui-react';
import { Story } from './reducers/stories';

type StoryListProps = {
  stories: Story[],
};

const StoryList: React.SFC<StoryListProps> = (props) => {
  return (
    <Item.Group>
      {props.stories.map((story, i) =>
        <Item key={i}>
          <Item.Content>
            <Item.Header as="a">{story.title}</Item.Header>
            <Item.Meta>By: {story.by}</Item.Meta>
            <Item.Extra>{story.score} {story.time}</Item.Extra>
          </Item.Content>
        </Item>
      )}
    </Item.Group>
  );
};

export default StoryList;
